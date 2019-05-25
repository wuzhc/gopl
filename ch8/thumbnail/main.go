package main

import (
	"bufio"
	"gopl/ch8/thumbnail/thumbnail"
	"log"
	"os"
)

func main() {
	// 读取console文件名
	// 先创建一个scanner
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		file := scanner.Text()
		thumb, err := thumbnail.ImageFile(file)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(thumb)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}
}
