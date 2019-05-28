package main

import (
	"bufio"
	"flag"
	"fmt"
	"gopl/me/ftp/funcc"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var hFlag = flag.String("h", "127.0.0.1", "please input host\n")
var pFlag = flag.String("p", "9999", "please input port")

func main() {
	flag.Parse()
	addr := *hFlag + ":" + *pFlag
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	} else {
		defer conn.Close()
	}

	go func() {
		for {
			var buf = make([]byte, 1024)
			conn.Read(buf)

			res := string(funcc.GetValidByte(buf))
			if res == "ready upload" {
				f, _ := os.Open("/data/wwwroot/go/src/gopl/me/dfs/main.go")
				io.Copy(conn, f)
				f.Close()
			} else {
				fmt.Printf("xxxxx", res)
			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		command := scanner.Text()
		cmd, opt := funcc.ParseCommand(command)
		switch cmd {
		case "list":
		case "upload":
			sendServ(conn, cmd, opt)
		}
	}
}

func sendServ(conn net.Conn, cmd string, opt []string) {
	optStr := strings.Join(opt, " ")
	msg := cmd + " " + optStr
	conn.Write([]byte(msg))
}
