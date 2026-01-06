package main

import (
	"bufio"
	"net"
	"testing"
	"time"
)

func TestPing(t *testing.T) {
	go startServer(":6379")
	time.Sleep(100 * time.Millisecond)

	conn, err := net.Dial("tcp", ":6379")

	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("PING\r\n"))

	if err != nil {
		t.Error(err)
	}

	reader := bufio.NewReader(conn)

	resp, err := reader.ReadString('\n')

	if err != nil {
		t.Error(err)
	}

	if resp != "+PONG\r\n" {
		t.Errorf("resp is %s", resp)
	} else {
		t.Log(resp)
	}

}

func TestUnknownCommand(t *testing.T) {
	go startServer(":6380")
	time.Sleep(100 * time.Millisecond)

	conn, err := net.Dial("tcp", ":6380")

	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("UNKNOWN\r\n"))

	if err != nil {
		t.Error(err)
	}

	reader := bufio.NewReader(conn)

	resp, err := reader.ReadString('\n')

	if err != nil {
		t.Error(err)
	}

	if resp != "-ERR unknown command\r\n" {
		t.Errorf("resp is %s", resp)
	} else {
		t.Log(resp)
	}

}
