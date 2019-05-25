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

var vFlag = flag.Bool("v", false, "show detail")

func main() {
	flag.Parse()

	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// 通道用来接收文件大小
	filesizes := make(chan int64)
	// 同步等待goroutine
	var wg sync.WaitGroup

	for _, dir := range roots {
		wg.Add(1)
		go walkDir(dir, &wg, filesizes)
	}

	go func() {
		wg.Wait()
		close(filesizes)
	}()

	var nfiles, nbytes int64
	var tick <-chan time.Time
	if *vFlag {
		tick = time.Tick(500 * time.Millisecond)
	}

loop:
	for {
		select {
		case size, ok := <-filesizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}

	printDiskUsage(nfiles, nbytes)
}

// 打印硬盘使用率
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// 遍历目录
func walkDir(dir string, wg *sync.WaitGroup, filesizes chan<- int64) {
	defer wg.Done()

	for _, f := range dirents(dir) {
		if f.IsDir() {
			wg.Add(1)
			subDir := filepath.Join(dir, f.Name())
			go walkDir(subDir, wg, filesizes)
		} else {
			filesizes <- f.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

// 读取目录
// 限制同时最多只能20个并发
func dirents(dir string) []os.FileInfo {
	defer func() {
		<-sema
	}()

	sema <- struct{}{}

	fileInfo, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		return fileInfo
	}
}
