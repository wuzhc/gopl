package crawler

import (
	"fmt"
	"gopl/ch5/links2"
	"gopl/ch8/findlinks/downloader"
	"net/http"

	"golang.org/x/net/html"
)

// 爬虫一个页面
// 下载页面内容
// 解析页面中所有子链接,返回主程序以便继续爬取
func Do(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("打开%s链接失败,错误原因为%s\n", url, err)
	} else {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("打开%s页面失败,错误码为:%v", url, resp.StatusCode)
	}

	err = downloader.SavePaper(url, resp.Body)
	if err != nil {
		return nil, err
	}

	links, _ := links2.Extract(url)
	return links, nil
}

func foreachNode(n *html.Node, pre func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	// 查询子节点
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Data == "body" {
			fmt.Println(c.FirstChild, "xxxxxxxxxx")
		}
		foreachNode(c, pre)
	}
}
