// 目标:写出一个支持高并发的爬虫程序
// 广度搜索算法
package main

import (
	"gopl/ch8/findlinks/crawler"
	"log"
)

func main() {
	var seen = make(map[string]bool)
	var worklist = make(chan []string)
	var ready = make(chan string)
	var token = make(chan struct{}, 20)

	// 为什么要用goroutine,因为在同个goroutine中不能使用通道通信
	go func() {
		url := "http://127.0.0.1:9501/"
		worklist <- []string{url}
	}()

	// ready会阻塞,创建多个ready通道
	for i := 0; i <= 30; i++ {
		go func() {
			for url := range ready {
				// 利用token保证同时20个进行
				token <- struct{}{}
				newUrls, err := crawler.Do(url)
				<-token

				if err != nil {
					log.Println(err)
					continue
				}

				//worklist没有被消费完,继续写入数据被阻塞了,所以用协程方式去处理
				worklist <- newUrls
			}
		}()
	}

	for list := range worklist {
		for _, url := range list {
			if !seen[url] {
				seen[url] = true
				// 用goroutine写入ready通道,避免worklist阻塞
				go func(url string) {
					ready <- url
				}(url)
			}
		}
	}

	log.Println("all done")
}
