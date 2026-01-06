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

func TestSet(t *testing.T) {
	go startServer(":6381")
	time.Sleep(100 * time.Millisecond)

	conn, err := net.Dial("tcp", ":6381")

	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("SET mykey hello\r\n"))
	// _, err = conn.Write([]byte("GET key\r\n"))

	if err != nil {
		t.Error(err)
	}

	reader := bufio.NewReader(conn)

	resp, err := reader.ReadString('\n')

	if err != nil {
		t.Error(err)
	}

	if resp != "+OK\r\n" {
		t.Errorf("resp is %s", resp)
	} else {
		t.Log(resp)
	}
}

func TestSetGet(t *testing.T) {
	go startServer(":6382")
	time.Sleep(100 * time.Millisecond)

	conn, err := net.Dial("tcp", ":6382")

	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("SET mykey hello\r\n"))

	if err != nil {
		t.Error(err)
	}

	reader := bufio.NewReader(conn)

	resp, err := reader.ReadString('\n')

	if err != nil {
		t.Error(err)
	}

	if resp != "+OK\r\n" {
		t.Errorf("resp is %s", resp)
	} else {
		t.Log(resp)
	}

	conn, err = net.Dial("tcp", ":6382")
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("GET mykey\r\n"))

	if err != nil {
		t.Error(err)
	}

	reader = bufio.NewReader(conn)

	lengthLine, err := reader.ReadString('\n')

	if err != nil {
		t.Error(err)
	}

	if lengthLine != "$5\r\n" {
		t.Errorf("resp is %s", lengthLine)
	} else {
		t.Log(lengthLine)
	}

	valueLine, err := reader.ReadString('\n')

	if err != nil {
		t.Error(err)
	}

	if valueLine != "hello\r\n" {
		t.Errorf("resp is %s", valueLine)
	} else {
		t.Log(valueLine)
	}
}
