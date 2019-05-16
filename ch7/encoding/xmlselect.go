// 解析一个xml放到一个树中,树中的节点主要两种类型,一种是标签(包含子节点数组),一种是内容
package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"strings"
)

type Node interface {
	String() string
}
type CharData string

func (c CharData) String() string {
	return string(c)
}

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e *Element) String() string {
	var attrs, children string
	for _, attr := range e.Attr {
		attrs += fmt.Sprintf(" %s=%q", attr.Name.Local, attr.Value)
	}
	for _, child := range e.Children {
		children += child.String()
	}
	return fmt.Sprintf("<%s %s>%s</%s>", e.Type.Local, attrs, children, e.Type.Local)
}

func Parse(input string) (Node, error) {
	inputReader := strings.NewReader(input)
	decoder := xml.NewDecoder(inputReader)
	var stack []*Element
	for t, err := decoder.Token(); err != io.EOF; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			// 入栈
			elem := &Element{Type: token.Name, Attr: token.Attr, Children: []Node{}}
			if len(stack) > 0 {
				lelem := stack[len(stack)-1]
				lelem.Children = append(lelem.Children, elem)
			}
			stack = append(stack, elem)
		case xml.EndElement:
			// 出栈
			llen := len(stack)
			if llen == 0 {
				return nil, errors.New("unexpected tag closing")
			} else if llen == 1 {
				return stack[0], nil
			}
			stack = stack[:llen-1]
		case xml.CharData:
			// 内容
			if len(stack) > 0 {
				lelem := stack[len(stack)-1]
				lelem.Children = append(lelem.Children, CharData(token))
			}
		}
	}
	return nil, nil
}

func main() {
	input := `<Person id="name" value="eeee"><FirstName>Xu</FirstName><LastName>Xinhua</LastName></Person>`
	node, err := Parse(input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(node)
}
