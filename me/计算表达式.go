// 算数表达式算法

package main

import (
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

type Expr struct {
	str   string
	queue []string
	kuo   int
}

// 解析
func (e *Expr) Parse() (err error) {
	defer func() {
		if err2 := recover(); err2 != nil {
			err = fmt.Errorf("%v", err2)
		}
	}()

	if len(e.str) == 0 {
		panic("表达式不能为空")
	}

	var temp []rune
	var s scanner.Scanner
	s.Init(strings.NewReader(e.str))
	s.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanInts
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch tok {
		case scanner.Float, scanner.Int:
			e.queue = append(e.queue, s.TokenText())
		case '+', '-':
			if len(e.queue) == 0 {
				panic("错误表达式" + e.str)
			}

			if e.kuo == 0 {
				var tlen = len(temp) - 1
				for i := tlen; i >= 0; i-- {
					// 优先级小于等于,则top出栈后入队列
					e.queue = append(e.queue, string(temp[i]))
					temp = temp[:i]
				}
			}
			temp = append(temp, tok)
		case '*', '/':
			// 优先级高,直接入栈
			// 取最后一个判断
			if len(e.queue) == 0 {
				panic("错误表达式" + e.str)
			}

			if e.kuo == 0 {
				var tlen = len(temp)
				if tlen > 0 {
					last := temp[tlen-1]
					if last == '*' || last == '/' {
						for i := tlen - 1; i >= 0; i-- {
							// top出栈后入队列
							if temp[i] == '*' || temp[i] == '/' {
								e.queue = append(e.queue, string(temp[i]))
								temp = temp[:i]
							} else {
								break
							}
						}
					}
				}
			}
			temp = append(temp, tok)
		case '(':
			temp = append(temp, tok)
			e.kuo = e.kuo + 1
		case ')':
			var tlen = len(temp)
			var i int
			for i = tlen - 1; i >= 0; i-- {
				// top出栈后入队列
				if temp[i] == '(' {
					break
				} else {
					e.queue = append(e.queue, string(temp[i]))
				}
			}
			temp = temp[:i]
			e.kuo = e.kuo - 1
		}
	}

	if len(temp) > 0 {
		for n := len(temp) - 1; n >= 0; n-- {
			e.queue = append(e.queue, string(temp[n]))
		}
	}

	return nil
}

// 中缀计算
func (e *Expr) Eval() (result float64, err error) {
	defer func() {
		if err2 := recover(); err2 != nil {
			err = fmt.Errorf("%v", err2)
		}
	}()

	parseErr := e.Parse()
	if parseErr != nil {
		fmt.Println(parseErr)
	} else {
		fmt.Println("解析后为:", e)
	}

	if len(e.queue) <= 0 {
		result = float64(0)
		return
	}

	var temp []float64
	for _, v := range e.queue {
		if v == "+" || v == "*" || v == "-" || v == "/" {
			tlen := len(temp)
			if tlen < 2 {
				panic("表达式解析错误")
			}
			switch v {
			case "+":
				result = temp[tlen-1] + temp[tlen-2]
			case "*":
				result = temp[tlen-1] * temp[tlen-2]
			case "-":
				result = temp[tlen-2] - temp[tlen-1]
			case "/":
				result = temp[tlen-2] / temp[tlen-1]
			}
			temp = temp[:tlen-2]
			temp = append(temp, result)
		} else {
			vv, _ := strconv.ParseFloat(v, 64)
			temp = append(temp, vv)
		}
	}

	return
}

func main() {
	var expr = &Expr{
		str: "2*(1+1*(1+1*2))/4-(2+2)+45/9",
		kuo: 0,
	}

	result, err := expr.Eval()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("表达式:", expr.str, "=", result)
	}
}
