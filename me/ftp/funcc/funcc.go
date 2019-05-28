package funcc

import (
	"strings"
)

// 解析命令行
func ParseCommand(command string) (cmd string, opt []string) {
	var temp []string
	arr := strings.Split(command, " ")
	for _, v := range arr {
		if len(v) == 0 {
			continue
		}
		v = strings.TrimSpace(v)
		temp = append(temp, v)
	}

	if len(temp) == 0 {
		return "", nil
	}

	return temp[0], temp[1:]
}

func GetValidByte(src []byte) []byte {
	var str_buf []byte
	for _, v := range src {
		if v != 0 {
			str_buf = append(str_buf, v)
		}
	}
	return str_buf
}
