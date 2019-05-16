package main

import (
	"bufio"
	"errors"
	"fmt"
	"strings"
)

func main() {
	counts := make(map[string]int)
	input := "wuzhc tangjf pengjc ymg wuzhc"
	scanner := bufio.NewScanner(strings.NewReader(input))

	// scanner.Split(bufio.ScanWords)
	splitFunc := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		fmt.Printf("%t\t%d\t%s\n", atEOF, len(data), data)
		if atEOF {
			return 0, nil, errors.New("at eof")
		} else {
			return 0, nil, nil
		}
	}
	scanner.Split(splitFunc)

	// 设置缓冲区的大小
	buf := make([]byte, 2)
	scanner.Buffer(buf, bufio.MaxScanTokenSize)

	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	// 报错处理
	if scanner.Err() != nil {
		fmt.Printf("error:%s\n", scanner.Err())
	}

	for k, n := range counts {
		fmt.Println(k, ":", n)
	}

}
