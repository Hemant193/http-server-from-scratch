package main

import (
	"net"
	"log"
)

func main()  {
	
	// Create a listener
	connection, error := net.Listen("tcp", ":8080")

	if error != nil {
		log.Fatalf("Failed to start server at port: %v", error)
	}
	defer connection.Close()
	log.Println("Server started at port 8080")

	// Looping and accepting connections
	for {
		client, err := connection.Accept()
		if err != nil {
			log.Printf("Failed to accept connections: %v", err)
		}
	
		response := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 12\r\n\r\nHello, World!"
		_, err = client.Write([]byte(response))
		if err != nil {
			log.Printf("Failed to write response: %v", err)
		}
		client.Close()
	}

}