package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Minute())

	t = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	t = t.AddDate(0, 1, 0)
	fmt.Println(t.Format("2006-01-02 15:04:05"))
	fmt.Println(time.January)
}
