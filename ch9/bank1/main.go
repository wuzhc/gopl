// 限制一个goroutine操作共享变量
package main

import (
	"fmt"
	"sync"
)

var deposits = make(chan int) // 存款
var balances = make(chan int) // 查看余额

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

// 用于消费
type draw struct {
	amount int
	succed chan bool
}

var withdraw = make(chan draw)

func Withdraw(amount int) bool {
	ch := make(chan bool)
	withdraw <- draw{amount, ch}
	return <-ch
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case draw := <-withdraw:
			if draw.amount <= balance {
				draw.succed <- true
				balance -= draw.amount
			} else {
				draw.succed <- false
			}
		}
	}
}

func main() {
	go teller()

	// 设置存款为100
	Deposit(100)

	var res bool
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(amount int) {
			res = Withdraw(i)
			if !res {
				fmt.Printf("取出%v失败 \n", amount)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

	balance := Balance()
	fmt.Printf("balance = %v\n", balance)
}
