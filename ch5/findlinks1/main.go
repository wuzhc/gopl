package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// 递归解析html文档,提取所有的url链接
func main() {
	var links []string

	f, err := os.Open("test.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 从标准输出获取html
	doc, err := html.Parse(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 不能用NextSibling,因为文档的html是没有下一个的,所以需要递归
	for _, link := range visit(doc, links) {
		fmt.Println(link)
	}
}

func visit(n *html.Node, links []string) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, v := range n.Attr {
			if v.Key == "href" {
				links = append(links, v.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(c, links)
	}

	return links
}
