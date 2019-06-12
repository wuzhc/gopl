// /data/wwwroot/doc/zcnote目录生成脚本
// go run main.go
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"strings"
	"text/template"
)

const templ = `{{range .}}{{.}}{{end}}`

var scanDir = flag.String("dir", "/data/wwwroot/doc/zcnote", "Please input scan dir")

var filters = []string{
	"/data/wwwroot/doc/zcnote/zhongzhi",
	"/data/wwwroot/doc/zcnote/temp",
	"/data/wwwroot/doc/zcnote/interview",
	"/data/wwwroot/doc/zcnote/code",
	"/data/wwwroot/doc/zcnote/daily",
	"/data/wwwroot/doc/zcnote/images",
	"/data/wwwroot/doc/zcnote/test.php",
	"/data/wwwroot/doc/zcnote/README.md",
	"/data/wwwroot/doc/zcnote/SUMMARY.md",
	"/data/wwwroot/doc/zcnote/book.json",
}

func main() {
	flag.Parse()

	if *scanDir == "" {
		errExit("Please input scan dir")
	}

	items := readDir(*scanDir, 0)
	// errExit(items)
	if len(items) == 0 {
		errExit("No file")
	}

	in, err := os.OpenFile("/data/wwwroot/doc/zcnote/README.md", os.O_RDWR, 0)
	if err != nil {
		in.Close()
		errExit(err)
	}
	defer in.Close()

	in.WriteString("## 目录\n")
	report := template.Must(template.New("hellow").Parse(templ))
	if err := report.Execute(in, items); err != nil {
		errExit(err)
	}

	writeInfo(in)
}

func writeInfo(in *os.File) {
	in.WriteString("## 目录自动生成器\n")
	in.WriteString("- [https://github.com/wuzhc/gopl/blob/master/me/automatic-summary-generate/main.go](https://github.com/wuzhc/gopl/blob/master/me/automatic-summary-generate/main.go)\n")
}

// 读取目录
func readDir(dir string, depath int) []string {
	fileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		errExit(err)
	}

	var files []string
	for _, f := range fileInfo {
		name := f.Name()

		// Use absolute path to filter specified directories or files
		abPath := path.Join(dir, name)
		if inFilter(abPath) {
			continue
		}

		// Filter hidden directories or files, such as ".drea"
		if strings.HasPrefix(name, ".") {
			continue
		}

		blankPrefix := getBlankPrefix(depath)

		if f.IsDir() {
			record := fmt.Sprintf("%s* %s\n", blankPrefix, name)
			files = append(files, record)
			subDir := path.Join(dir, name)
			dh := depath + 1
			files = append(files, readDir(subDir, dh)...)
		} else {
			relativePath := getRelativePath(dir, name)
			labelName := getLableName(name)
			record := fmt.Sprintf("%s* [%s](%s)\n", blankPrefix, labelName, relativePath)
			files = append(files, record)
		}
	}

	return files
}

func getLableName(filename string) string {
	ext := path.Ext(filename)
	return strings.Replace(filename, ext, "", len(ext))
}

// 相对连接
func getRelativePath(dir string, name string) string {
	if dir == *scanDir {
		return ""
	}

	link := strings.Replace(dir, *scanDir+"/", "", len(*scanDir)) + "/" + name
	return url.PathEscape(link)
}

// 空白
func getBlankPrefix(depath int) string {
	return strings.Repeat("    ", depath)
}

// 指定过滤路径
func inFilter(f string) bool {
	for _, v := range filters {
		if v == f {
			return true
		}
	}

	return false
}

// 错误退出
func errExit(msg ...interface{}) {
	var errMsg string
	if len(msg) > 0 {
		s := strings.Repeat("%v ", len(msg))
		errMsg = fmt.Sprintf(s, msg...)
	} else {
		errMsg = "Unkown"
	}

	fmt.Println(errMsg)
	os.Exit(1)
}
