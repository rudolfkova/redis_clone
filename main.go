package main

import (
	"bufio"
	"log"
	"net"
	redis "redis/internal"
	"strings"
)

func main() {
	startServer(":6379")
}

func startServer(addr string) {
	storage := redis.NewStorage()
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

		go func(conn net.Conn, storage redis.Storage) {
			defer conn.Close()
			reader := bufio.NewReader(conn)
			req, err := reader.ReadString('\n')
			if err != nil {
				log.Print(err)
				return
			}
			req = strings.TrimRight(req, "\r\n")
			reqSlice := strings.Fields(req)

			switch reqSlice[0] {
			case "PING":
				err = Ping(req, conn)
			case "SET":
				err = Set(reqSlice, conn, storage)
			case "GET":
				err = Get(reqSlice, conn, storage)

			default:
				UnknownCommand(conn)
			}

			if err != nil {
				log.Print(err)
			}

		}(conn, storage)
	}
}
