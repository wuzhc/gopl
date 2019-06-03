package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		done := make(chan int)
		go func(v int) {
			defer wg.Done()
			time.Sleep(4 * time.Second)
			done <- v
		}(i)

		select {
		case <-time.After(3 * time.Second):
			go func() {
				v := <-done // 防止所有goroutine都退出了,而done通道还有数据没消费完
				log.Println("xxxx", v)
			}()
			log.Println("time out", i)
		case v := <-done:
			log.Println("get ", v)
		}
	}

	wg.Wait()
	log.Println("end..........")
}
