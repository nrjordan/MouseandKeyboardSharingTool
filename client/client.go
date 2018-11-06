package client

import (
	"bufio"
	"fmt"
	"mkShareReceiver/mkLogger"
	"net"
)

var logger = mkLogger.GetInstance()

type client struct {
	*bufio.Reader
	*bufio.Writer
	wc chan string
}

func StartClient(ipAddress string) {

	conn, err := net.Dial("tcp", ipAddress+":8989")
	if err != nil{
		logger.Println(err)
		fmt.Println("Error connecting to IP...")
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}

func (c *client) writeMonitor() {
	go func() {
		for s := range c.wc {
			logger.Println("Sending", s)
			c.WriteString(s + "\n")
			c.Flush()
		}
	}()
}