package main

import (
	"fmt"
	"gopl/ch4/github"
	"os"
	"strings"
	"text/template"
)

// 定义一个模板
const templ = `{{.TotalCount}} issues:
{{range .Items}}-----------------------
Title: {{.Title | printf "%.64s"}}
Url: {{.Url}}
User: {{.User.Login | formatUser}}
{{end}}
`

func formatUser(u string) string {
	return "u_" + u
}

// go run main.go windows label:bug language:python state:open
func main() {
	result, _ := github.SearchIssues(os.Args[1:])
	if result == nil {
		fmt.Println("no result with: ", strings.Join(os.Args[1:], ","))
		os.Exit(1)
	}

	report := template.Must(template.New("hellow").
		Funcs(template.FuncMap{"formatUser": formatUser}).
		Parse(templ))

	if err := report.Execute(os.Stdout, result); err != nil {
		fmt.Println(err)
	}
}
