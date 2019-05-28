// 多个目录分别统计
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type sizeInfo struct {
	id   int
	size int64
}

var vFlag = flag.Bool("v", false, "show detail")

func main() {
	flag.Parse()

	// 是否定时显示
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(1 * time.Second)
	}

	// 获取命令行需要统计的目录slice
	roots := flag.Args()
	if len(roots) == 0 {
		// 默认为当前目录
		roots = []string{"."}
	}

	var wg sync.WaitGroup
	var ch = make(chan sizeInfo)

	for id, dir := range roots {
		wg.Add(1)
		go walkDir(dir, &wg, id, ch)
	}

	// 等待所有goroutine退出
	go func() {
		wg.Wait()
		close(ch)
	}()

	// 统计
	var nfiles = make(map[int]int, len(roots))
	var nbytes = make(map[int]int64, len(roots))
loop:
	for {
		select {
		case res, ok := <-ch:
			if !ok {
				break loop
			}
			nfiles[res.id]++
			nbytes[res.id] += res.size
		case <-tick:
			printDickUsage(nfiles, nbytes, roots)
		}
	}

	printDickUsage(nfiles, nbytes, roots)
}

// 打印硬盘使用状况
func printDickUsage(nfiles map[int]int, nbytes map[int]int64, roots []string) {
	for id, _ := range roots {
		fmt.Printf("dir:%s files: %v bytes: %.1fG\n", roots[id], nfiles[id], float64(nbytes[id])/1e9)
	}
}

func walkDir(dir string, wg *sync.WaitGroup, id int, ch chan sizeInfo) {
	defer wg.Done()
	for _, f := range direvents(dir) {
		if f.IsDir() {
			subDir := filepath.Join(dir, f.Name())
			wg.Add(1)
			go walkDir(subDir, wg, id, ch)
		} else {
			ch <- sizeInfo{id, f.Size()}
		}
	}
}

// 限制并发度,防止打开太多文件
var openLimit = make(chan struct{}, 20)

func direvents(dir string) []os.FileInfo {
	defer func() { <-openLimit }()
	openLimit <- struct{}{}

	fileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return fileInfo
}
