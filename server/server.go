package server

import (
	"bufio"
	"fmt"
	"mkShareReceiver/mkLogger"
	"net"
)

//var transmit bool
var logger = mkLogger.GetInstance()
var ln net.Listener
var err error

func StartServer(res float64) error {
	ln, err = net.Listen("tcp", ":8989")
	if err != nil {
		logger.Println("Error setting up the listener: ", err)
		return err
	}
	logger.Println("Listener started...", res)
	fmt.Println("Listener started", res)
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				logger.Println("Problem accepting connection: ", err)
				fmt.Println(err)
				break
			}
			handleConnection(conn)
		}
	}()

	return err
}

func StopServer() {
	err := ln.Close()
	if err != nil {
		logger.Println("Error closing connection: ", err)
		fmt.Println(err)
	}
	logger.Println("Connection closed")
	fmt.Println("Connection closed.")
}

func handleConnection(conn net.Conn) {
	fmt.Println("Received connection!")
	logger.Println("Received connection from", conn.RemoteAddr())
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}