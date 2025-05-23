/*
 * Simple HTTP Server
 * 
 * A basic HTTP server implementation using raw TCP connections in Go.
 * This server listens on port 8080 and responds with "Hello, World!" 
 * to all incoming HTTP requests.
 * 
 * Author: Hemant Mishra
 * Date: May 23, 2025
 
 */

package main

import (
	"net"
	"log"
)

// Server configuration constants
const (
	// PORT defines the port number the server listens on
	PORT = ":8080"
	
	// HTTP_RESPONSE is the complete HTTP response sent to clients
	HTTP_RESPONSE = "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 12\r\n\r\nHello, World!"
)


// main initializes and starts the HTTP server
func main()  {
	// Create a TCP listener on the specified port
	connection, error := net.Listen("tcp", PORT)
	if error != nil {
		log.Fatalf("Failed to start server at port: %v", error)
	}
	defer connection.Close()

	log.Println("Server started at port 8080")

	// Main server loop - accept and handle incoming connections
	for {
		client, err := connection.Accept()
		if err != nil {
			log.Printf("Failed to accept connections: %v", err)
			continue // To accept new connections
		}

		// Handle each connection concurrently in a separate goroutine
		go handleReq(client)
	}
}

// HandleConnection processes an individual client connection
// It sends an HTTP response and closes the connection
func handleReq(conn net.Conn) {
		
	// Send HTTP response to the client
	_, err := conn.Write([]byte(HTTP_RESPONSE))
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}
	// Ensure the connection is closed when function exits
	conn.Close()
}