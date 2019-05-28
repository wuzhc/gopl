// 类似于ftp功能一样
// 先进入,等待命令
// ls dir
// get file
// send file
// rm file
package main

import (
	"flag"
	"gopl/me/ftp/funcc"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

// var hFlag = flag.String("h", "10.8.8.185", "please input host")
var hFlag = flag.String("h", "127.0.0.1", "please input host")
var pFlag = flag.String("p", "9999", "please input port")

func main() {
	// 监听端口
	addr := *hFlag + ":" + *pFlag
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	} else {
		defer listener.Close()
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn)
	}

	log.Println("end...")
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		conn.Write([]byte("ftp: "))
		var buf = make([]byte, 1024)
		conn.Read(buf)
		cmd, opt := funcc.ParseCommand(string(funcc.GetValidByte(buf)))
		log.Println(cmd, opt)
		switch cmd {
		case "list":
			listDir(conn, opt)
		case "read":
		case "download":
		case "upload":
			conn.Write([]byte("ready upload"))
			f, _ := os.Create("/data/wwwroot/go/src/gopl/me/dfs/test_main.php")
			io.Copy(f, conn)
			f.Close()
			go func() {
				conn.Write([]byte("has upload...\n"))
			}()
		}
	}
}

func listDir(conn net.Conn, opt []string) {
	if len(opt) == 0 {
		conn.Write([]byte("please input list target directory"))
		return
	}

	var dirs []string
	for _, dir := range opt {
		fileInfo, err := ioutil.ReadDir(string(dir))
		if err != nil {
			conn.Write([]byte(err.Error()))
			continue
		}

		for _, f := range fileInfo {
			dirs = append(dirs, f.Name())
		}
	}

	str := strings.Join(dirs, " ")
	str += "\n"
	conn.Write([]byte(str))
}
