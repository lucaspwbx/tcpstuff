package main

import (
	"fmt"
	"net"
	"os"
)

const (
	uiDir  = "dir"
	uiCd   = "cd"
	uiPwd  = "pwd"
	uiQuit = "quit"
)

const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1202")
	checkError(err)
	for {
		cdRequest(conn, "teste")
		buf := make([]byte, 512)
		n, _ := conn.Read(buf[0:])
		s := string(buf[0:n])
		fmt.Println(s)
	}
}

func dirRequest(conn net.Conn) {
	conn.Write([]byte(DIR + " "))
}

func cdRequest(conn net.Conn, dir string) {
	conn.Write([]byte(CD + " "))
}

func pwdRequest(conn net.Conn) {
	conn.Write([]byte(PWD + " "))
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
