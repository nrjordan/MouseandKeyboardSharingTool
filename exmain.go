package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8989")
	if err != nil {
		fmt.Println("Error setting up listen")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("Received connection!")
	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	/*keve := robotgo.AddEvent("k")
	if keve == 0 {
		fmt.Println("you press...", "k")
	}

	mleft := robotgo.AddEvent("mleft")
	if mleft == 0 {
		fmt.Println("you press...", "mouse left button")
	}*/
}
