package server

import (
	"bufio"
	"encoding/gob"
	"errors"
	"github.com/go-vgo/robotgo"
	"io"
	"mkShareReceiver/mkLogger"
	"net"
	"strings"
	"sync"
)

type mouse struct {
	x, y  int
	click chan bool
}

//var transmit bool
var logger = mkLogger.GetInstance()
var ln net.Listener
var err error

/*func StartServer() error {
	ln, err = net.Listen("tcp", ":8989")
	if err != nil {
		logger.Println("Error setting up the listener: ", err)
		return err
	}
	logger.Println("Listener started...")
	fmt.Println("Listener started")
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
}*/

type HandleFunc func(*bufio.ReadWriter)

type Endpoint struct {
	listener net.Listener
	handler  map[string]HandleFunc

	m sync.RWMutex
}

func NewEndpoint() *Endpoint {
	return &Endpoint{
		handler: map[string]HandleFunc{},
	}
}

func (e *Endpoint) AddHandleFunc(name string, f HandleFunc) {
	e.m.Lock()
	e.handler[name] = f
	e.m.Unlock()
}

func (e *Endpoint) Listen() error {
	var err error
	e.listener, err = net.Listen("tcp", ":8989")
	if err != nil {
		return errors.New("unable to listen on port 8989")
	}
	logger.Println("Listening on", e.listener.Addr().String())
	for {
		logger.Println("Accepting connection request")
		conn, err := e.listener.Accept()
		if err != nil {
			logger.Println("Failed to accept connection", err)
			continue
		}
		logger.Println("Handling incoming messages")
		go e.handleMessages(conn)
	}
}

func (e *Endpoint) handleMessages(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	defer conn.Close()

	for {
		logger.Println("Received command")
		cmd, err := rw.ReadString('\n')
		switch {
		case err == io.EOF:
			logger.Println("Reached EOF - closing connection.")
			return
		case err != nil:
			logger.Println("Error reading command.", cmd, err)
			return
		}

		cmd = strings.Trim(cmd, "\n ")
		logger.Println(cmd + "'")
		e.m.RLock()
		handleCommand, ok := e.handler[cmd]
		e.m.RUnlock()
		if !ok {
			logger.Println("Command '" + cmd + "' failed")
			return
		}
		handleCommand(rw)
	}
}

func handleKey(rw *bufio.ReadWriter) {
	logger.Print("Received key press")
	k, err := rw.ReadString('\n')
	if err != nil {
		logger.Println("Unable to read key...", err)
	}
	k = strings.Trim(k, "\n ")
	robotgo.KeyTap(k)
	err = rw.Flush()
	if err != nil {
		logger.Println("Flush failed.", err)
	}
}

func handleGob(rw *bufio.ReadWriter) {
	logger.Println("Received mouse data")
	var data mouse
	dec := gob.NewDecoder(rw)
	err := dec.Decode(&data)
	if err != nil {
		logger.Println("Error decoding mouse data:", err)
		return
	}
	logger.Println("x data:", data.x)
	logger.Println("y data:", data.y)
}

func StartServer() error {
	endpoint := NewEndpoint()
	endpoint.AddHandleFunc("KEY", handleKey)
	endpoint.AddHandleFunc("GOB", handleGob)
	return endpoint.Listen()
}
