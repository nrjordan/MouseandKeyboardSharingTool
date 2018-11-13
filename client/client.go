package client

import (
	"bufio"
	"fmt"
	"mkShareReceiver/mkLogger"
	"net"
)

var conn net.Conn
var logger = mkLogger.GetInstance()

type client struct {
	*bufio.Reader
	*bufio.Writer
	wc chan string
}

type mouse struct {
	x, y  int
	click chan bool
}

func StartClient(ipAddress string) (string, error) {

	conn, err := net.Dial("tcp", ipAddress+":8989")
	if err != nil {
		logger.Println(err)
		return "Error connecting to IP...", err
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		logger.Println(err)
		return "Error with reader...", err
	}
	fmt.Println(status)
	return "Success!", nil
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

/*func sendMouse() {
	for {
		x, y := robotgo.GetMousePos()
		conn.Write(x, y)
	}
}

func createWindow() {
	if runtime.GOOS == "windows"{

	}
}*/
