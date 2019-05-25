package links2

import (
	"errors"
	"net/http"

	"golang.org/x/net/html"
)

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	} else {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("请求url失败")
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	visitNode := func(n *html.Node) {
		// 闭包,引用links变量
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, v := range n.Attr {
				if v.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(v.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}

	foreachNode(doc, visitNode)
	return links, nil
}

func foreachNode(n *html.Node, pre func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	// 查询子节点
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		foreachNode(c, pre)
	}
}
