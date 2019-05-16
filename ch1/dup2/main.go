package main

import (
	"bufio"
	"bytes"
	"errors"
	"os"
)

import (
	"fmt"
)

func main() {
	counts := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)

	// 设置终止符号
	splitFunc := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if bytes.Equal(token, []byte{'e', 'n', 'd'}) {
			return 0, []byte{'e', 'n', 'd'}, errors.New("at eof")
		}
		return
	}
	scanner.Split(splitFunc)

	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	for v, n := range counts {
		fmt.Println(v, "---", n)
	}
}
