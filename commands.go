package main

import (
	"fmt"
	"net"
	redis "redis/internal"
	"strings"
)

func UnknownCommand(conn net.Conn) error {
	_, err := conn.Write([]byte("-ERR unknown command\r\n"))
	return err

}

func Ping(req string, conn net.Conn) error {
	_, err := conn.Write([]byte("+PONG\r\n"))
	return err
}

func Set(reqSlice []string, conn net.Conn, storage redis.Storage) error {
	if len(reqSlice) < 3 {
		return fmt.Errorf("Miss key or value in SET command")
	}
	key := reqSlice[1]
	value := strings.Join(reqSlice[2:], " ")
	storage.Set(key, value)
	_, err := conn.Write([]byte("+OK\r\n"))
	return err
}

func Get(reqSlice []string, conn net.Conn, storage redis.Storage) error {
	if len(reqSlice) < 2 {
		return fmt.Errorf("Miss key in GET command")
	}
	key := reqSlice[1]
	value, ok := storage.Get(key)
	if !ok {
		_, err := conn.Write([]byte("$-1\r\n"))
		return err
	}
	_, err := conn.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(value), value)))
	return err
}
