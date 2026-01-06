package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func main() {
	startServer(":6379")
}

func startServer(addr string) {
	listener, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}

		go func(conn net.Conn) {
			defer conn.Close()
			reader := bufio.NewReader(conn)
			req, err := reader.ReadString('\n')
			if err != nil {
				log.Print(err)
				return
			}

			Ping(req, conn)

		}(conn)
	}
}

func Ping(req string, conn net.Conn) {
	req = strings.TrimRight(req, "\r\n")
	if strings.ToUpper(req) == "PING" {
		_, err := conn.Write([]byte("+PONG\r\n"))
		if err != nil {
			log.Print(err)
		}
	} else {
		_, err := conn.Write([]byte("-ERR unknown command\r\n"))
		if err != nil {
			log.Print(err)
		}
	}
}
