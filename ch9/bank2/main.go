// 互斥锁 sync.Mutex
package main

import (
	"fmt"
	"sync"
)

var (
	wg      sync.WaitGroup
	mu      sync.Mutex
	balance int
)

// 存款
func Deposit(amount int) {
	defer mu.Unlock()
	mu.Lock()
	deposit(amount)
}

// 余额
func Balance() int {
	defer mu.Unlock()
	mu.Lock()
	return balance
}

func Withdraw(amount int) bool {
	defer mu.Unlock()
	mu.Lock()

	if amount <= balance {
		balance -= amount
		return true
	} else {
		return false
	}
}

func deposit(amount int) {
	balance += amount
}

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(amount int) {
			Deposit(amount)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Printf("剩余%v\n", balance)
}
