package main

import (
	"fmt"
	"net"
	"os"
)

const (
	DIR = "DIR"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
	service := "0.0.0.0:1202"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	fmt.Println("Listening on port 1202")
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		fmt.Println("Entrou uma vez")
		n, err := conn.Read(buf[0:])
		//	fmt.Println(n)
		if err != nil {
			conn.Close()
			return
		}
		s := string(buf[0:n])
		fmt.Println(len(s))
		//	fmt.Println(s[0:2])
		if s[0:2] == CD {
			fmt.Println("CD requested")
			conn.Write([]byte("OK"))
		} else if s[0:3] == DIR {
			fmt.Println("Dir requested")
		} else if s[0:3] == PWD {
			fmt.Println("PWD requested")
		} else {
			fmt.Println("Fail")
		}
	}
}

func chdir(conn net.Conn, s string) {
	if os.Chdir(s) == nil {
		conn.Write([]byte("OK"))
	} else {
		conn.Write([]byte("ERROR"))
	}
}

func pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		conn.Write([]byte(""))
		return
	}
	conn.Write([]byte(s))
}

func dirList(conn net.Conn) {
	defer conn.Write([]byte("\r\n"))

	dir, err := os.Open(".")
	if err != nil {
		return
	}

	names, err := dir.Readdirnames(-1)
	if err != nil {
		return
	}
	for _, nm := range names {
		conn.Write([]byte(nm + "\r\n"))
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
