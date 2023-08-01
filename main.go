package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const address = "localhost:8080"

func main() {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Can't create listener: ", err)
	}

	go func() {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal("Error listener accept: ", err)
		}
		defer connection.Close()

		connection.SetDeadline(time.Now().Add(1 * time.Second))

		// n - number of bytes the server sent
		n, err := connection.Write([]byte("Hello from sever ^_^"))
		if err != nil {
			log.Fatal("Error to write message: ", err)
		}
		log.Printf("Sended %v bytes...\n", n)
	}()

	connection, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal("Can't connect to TCP-address: ", err)
	}
	defer connection.Close()

	buffer := make([]byte, 0, 4096)
	tmp := make([]byte, 256)

	for {
		n, err := connection.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("End of File...")
			}
			break
		}
		log.Printf("Read all %v bytes...\n", n)
		buffer = append(buffer, tmp[:n]...)
	}

	log.Printf("Read all %v bytes...", len(buffer))
	fmt.Printf("%s\n", buffer)
}
