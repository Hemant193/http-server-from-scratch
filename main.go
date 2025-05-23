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

}