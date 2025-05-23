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
		continue
	
		// Closing the connection
		client.Close()
	}

	

}