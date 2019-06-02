package main

import (
	"bytes"
	"log"
	"os/exec"
)

func main() {
	err, stdout, stderr := ShellExec("echo \"hello world\" >> text.txt")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("stdout------------", stdout)
	log.Println("stderr------------", stderr)
}

// sh -c "command"
// e.g. sh -c "echo "hello world" >> text.txt"
// sh -c 可以让bash将一个字符串作为完整命令来执行
func ShellExec(command string) (error, string, string) {
	var stdout, stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}
