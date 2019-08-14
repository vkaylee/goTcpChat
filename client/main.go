package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:2349")
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()
	reader := bufio.NewReader(os.Stdin)
	inputName, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	for {
		inputText, err := reader.ReadString('\n')
		if err != nil {
			log.Println(err)
			break
		}
		msg := fmt.Sprintf("%s: %s\n", inputName[:len(inputName)-1], inputText[:len(inputText)-1])
		_, err = connection.Write([]byte(msg))
		if err != nil {
			log.Fatal(err)
		}
	}
}
