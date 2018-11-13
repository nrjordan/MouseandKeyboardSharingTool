package main

import (
	"bufio"
	"fmt"
	"mkShareReceiver/client"
	"mkShareReceiver/server"
	"os"
	"strings"
)

func main() {
	q := false
	serv := false
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Mouse and Keyboard Sharing software package.\n\nPress 1 if this computer is to be " +
		"remotely accessed, or 2 if this computer will be accessing another. When you are finished, press 'q' to quit.")
	for q == false {
		choice, _ := reader.ReadString('\n')
		choice = strings.Replace(choice, "\n", "", -1)
		if choice == "1" && serv == false {
			fmt.Println("Setting up the listener...Use the 's' key to stop the server.")
			server.StartServer()
			serv = true
		} else if choice == "2" {
			fmt.Println("Please enter the IP address of the computer you wish to access.")
			ip, _ := reader.ReadString('\n')
			ip = strings.Replace(ip, "\n", "", -1)
			client.StartClient(ip)
		} else if choice == "q" {
			q = true
			/*} else if choice == "s" && serv == true {
			fmt.Println("Stopping server.")
			server.StopServer()
			serv = false*/
		} else {
			fmt.Println("Invalid selection. Please try again.")
		}
	}

}
