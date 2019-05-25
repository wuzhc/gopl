package main

import (
	"fmt"
	"gopl/ch4/github"
	"html/template"
	"os"
	"strings"
)

const templ = `<h1>{{.TotalCount}}</h1>issues
<table>
<th><td>标题</td><td>链接</td><td>作者</td></th>
{{range .Items}}
<tr><td>{{.Title}}</td><td>{{.Url}}</td><td>.User.Login</td><tr>
{{end}}
</table>
`

func main() {
	terms := os.Args[1:]
	result, err := github.SearchIssues(terms)
	if err != nil {
		fmt.Println("no result with: ", strings.Join(os.Args[1:], ","))
		os.Exit(1)
	}

	report := template.Must(template.New("issuehtml").
		Parse(templ))

	err = report.Execute(os.Stdout, result)
	if err != nil {
		fmt.Println("parse failed")
		os.Exit(1)
	}
}
