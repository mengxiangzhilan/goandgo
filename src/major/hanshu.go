package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

/**
形参是实参的拷贝，实参是值传递，对形参修改不会影响实参，但如果是引用类型就不一样了
递归使函数可以直接或间接调用自己
当函数所有返回值都显式指定时 return后不需要返回值

*/

func Hanshumain() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
