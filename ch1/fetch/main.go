package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		hasPrefix := strings.HasPrefix(url, "http")
		if hasPrefix == false {
			fmt.Fprintf(os.Stderr, "The url %v has not http prefix\n", url)
			os.Exit(1)
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		} else {
			fmt.Fprintf(os.Stdin, "http status is %v\n", resp.Status)
		}

		b, err := ioutil.ReadAll(resp.Body)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
