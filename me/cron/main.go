package main

import (
	"log"

	"github.com/robfig/cron"
)

func main() {
	ch := make(chan struct{})

	i := 0
	c := cron.New()
	spec := "*/1 * * * * ?"
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.AddFunc("@every 1h1m", func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()

	<-ch
}
