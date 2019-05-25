package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

var ch = make(chan int64)

func main() {
	var files []string

	go func() {
		for _, dir := range os.Args[1:] {
			files, _ = findFiles(dir)
			// for _, f := range files {
			// fmt.Println(f)
			// }
		}
		close(ch)
	}()

	var nfiles, nbytes int64
	for size := range ch {
		nfiles++
		nbytes += size
	}

	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func findFiles(dir string) ([]string, error) {
	var files []string
	info, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, f := range info {
		if f.IsDir() {
			subDir := filepath.Join(dir, f.Name())
			cf, _ := findFiles(subDir)
			files = append(files, cf...)
		} else {
			subFile := filepath.Join(dir, f.Name())
			files = append(files, subFile)
			ch <- f.Size()
			fmt.Println(subFile)
		}
	}

	return files, nil
}
