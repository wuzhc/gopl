package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func SavePaper(url string, in io.Reader) error {
	name := filepath.Base(url)
	ext := filepath.Ext(url)
	name = strings.Trim(name, ext)
	fname := "./html/" + name + ".html"

	f, err := os.Create(fname)
	if err != nil {
		return fmt.Errorf("创建%s失败\n", fname)
	} else {
		defer f.Close()
	}

	_, err = io.Copy(f, in)
	if err != nil {
		return fmt.Errorf("写入%s失败,错误原因:%v\n", fname, err)
	}

	return nil
}

func Do(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("下载%s链接失败,错误原因为%s\n", url, err)
	} else {
		defer resp.Body.Close()
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("获取%s页面失败,错误码为:%v", url, resp.StatusCode)
	}

	return SavePaper(url, resp.Body)
}
