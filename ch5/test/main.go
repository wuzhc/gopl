package main

import (
	"fmt"
	"gopl/ch5/links2"
	"os"
)

func main() {
	links, err := links2.Extract("https://github.com/wuzhc")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, link := range links {
		fmt.Println(link)
	}
}
