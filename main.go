package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

var (
	// clientMap will hold the client connections with their respective usernames.
	clientMap = make(map[string]net.Conn)
	mu        sync.Mutex // To synchronize access to clientMap
)

func handleConnection(conn net.Conn, username string) {
	defer conn.Close()
	// Add the client to the map.
	mu.Lock()
	clientMap[username] = conn
	mu.Unlock()

	// Send a welcome message to the client.
	_, err := conn.Write([]byte("Hello from Server!"))
	if err != nil {
		log.Printf("Error writing to connection: %v", err)
	}

	for {
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Printf("%s disconnected", username)
			} else {
				log.Printf("Error reading: %v", err)
			}
			// Remove the client from the map when they disconnect
			mu.Lock()
			delete(clientMap, username)
			mu.Unlock()
			return
		}

		received := string(buf[:n])
		log.Printf("Received from %s", received)

		// Forward the message to all other clients except the sender
		mu.Lock()
		for otherUsername, otherConn := range clientMap {
			if otherUsername != username { // Don't send to the sender
				_, writeErr := otherConn.Write([]byte(received))
				if writeErr != nil {
					log.Printf("Error sending to %s: %v", otherUsername, writeErr)
				}
			}
		}
		mu.Unlock()
	}
}

func startServer() error {
	server := ReadConfig() // Assuming this function reads your server config.
	addr := fmt.Sprintf("%s:%d", server.IP, server.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("error listening: %w", err)
	}
	defer listener.Close()

	log.Printf("%s Listening on %s\n", server.Name, addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Handle the initial message to get the username
		buf := make([]byte, 4096)
		n, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Println("Client disconnected")
			} else {
				log.Printf("Error reading: %v", err)
			}
			conn.Close()
			continue
		}

		username := strings.Split(string(buf[:n]), ":")[0]
		log.Printf("New client connected: %s from %s", username, conn.RemoteAddr().String())
		go handleConnection(conn, username)
	}
}

func main() {
	if err := startServer(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
