package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http") == false {
			fmt.Fprintf(os.Stdout, "url %v has not http\n", url)
			continue
		}
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		// chan是不能直接for range的,chan会阻塞,直到有数据返回
		fmt.Println(<-ch)
	}

	fmt.Printf("total custom %.2fs\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("%v:%v\n", url, err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("%v:%v\n", url, err)
		return
	}

	ch <- fmt.Sprintf("time:%vs bytes:%v url:%v\n", time.Since(start).Seconds(), nbytes, url)
}
