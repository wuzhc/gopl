package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func Hello() []string {
	return []string{"hello word"}
}

func Extract(url string) ([]string, error) {
	fmt.Println("begin get url: ", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("请求失败,错误为%v\n", resp.StatusCode)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	links = visitNode(doc, links)
	return links, nil
}

func visitNode(n *html.Node, links []string) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, v := range n.Attr {
			if v.Key == "href" {
				links = append(links, v.Val)
			}
		}
	}

	// 递归查询子节点
	// 从第一个子节点开始
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visitNode(c, links)
	}

	return links
}
