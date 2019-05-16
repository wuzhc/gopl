package main

import (
	"fmt"
	"gopl/ch4/github"
	"os"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("总共有%d记录\n", result.TotalCount)
	for _, v := range result.Items {
		fmt.Println(v.User.GetName())
	}
}
