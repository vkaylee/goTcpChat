package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var (
	conns         []net.Conn
	activeConnsCh = make(chan net.Conn)
	closeConnsCh  = make(chan net.Conn)
	messagesCh    = make(chan string)
)

func main() {
	// Declare port server
	port := "2349"
	// Listen on port
	server, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	// Disconnect server when closing
	defer server.Close()
	// Show on terminal the status
	fmt.Println("Chat server is listening on port ", server.Addr())
	go func() {
		// Create the loop forever
		for {
			// Accept every connection
			connection, err := server.Accept()
			if err != nil {
				log.Fatal(err)
			}
			// Show the connection on terminal
			fmt.Println("New client", connection)
			// append the connection to list connections
			conns = append(conns, connection)
			activeConnsCh <- connection
			// Print the number of connections
			fmt.Printf("Current connection: %d\n", len(conns))
		}
	}()
	for {
		select {
		case connection := <-activeConnsCh:
			go listenMessage(connection)
		case conn := <-closeConnsCh:
			go log.Println(conn, " exits")
		case msg := <-messagesCh:
			go fmt.Println(msg[:len(msg)-1])
		}
	}
}
func listenMessage(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		messagesCh <- msg
	}
	closeConnsCh <- conn
}
