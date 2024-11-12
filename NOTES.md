ERRORS & FIXES :

PROBLEM: Server was waiting indefinitely

Need to start Server_Session in a separate go routine instead of just invoking both client and server like this

Server_Session()
StartClient() 

both Server_Session() and StartClient() are being called sequentially, which means the client will only start after Server_Session() completes. Since Server_Session() is starting a TCP listener and then entering an indefinite loop to accept incoming connections, it doesn’t return or allow StartClient() to execute, which leads to the observed hang.

FIX :

  // Start the server in a separate goroutine
  go Server_Session()

  // Start the client session
  StartClient()

## Exporting Functions

To use functions even within the same package but different scripts, need to use uppercase for defining functions

## RPC Calls
There is a specific function signature that needs to be used for RPC calls

For RPC in Go, it's not enough for just the function names to be uppercase. The function signatures must also follow the:

- func (t *Type) MethodName(args *Type1, reply *Type2) error pattern:

Arguments and Reply Requirements: Each RPC method needs to accept exactly two arguments:
- A pointer to an argument struct (args *Type1).
- A pointer to a reply struct (reply *Type2).
- Error Return Type: The method must return an error


## Setting up server as a TCP listener and accept incoming RPC connections
To modify your server_session function to set up the server as a TCP listener and accept incoming RPC connections,
you can add the necessary net and rpc logic. Here’s an updated version of the server_session function that integrates the functionality for 
initializing seats, starting the request monitor, registering the ServerSession struct for RPC, and beginning to listen for incoming client connections.

### Steps
#### RPC Registration: 
rpc.Register(server) registers the methods of ServerSession for RPC, so clients can call these methods.
#### Listener Initialization: 
net.Listen("tcp", "127.0.0.1:8000") sets up a TCP listener on 127.0.0.1:8000.
#### Accepting Connections: 
A for loop is used to accept incoming connections, each handled in a new goroutine by rpc.ServeConn(conn).


package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

// could be main function
func server_session() {
	logger := log.Default()

	// Initialize seats globally and start monitoring requests
	InitSeats()
	go MonitorLockRequestRelease()

	// Register ServerSession methods with RPC so they can be called by the client
	server := new(ServerSession)
	err := rpc.Register(server)
	if err != nil {
		log.Fatal("Error registering ServerSession:", err)
	}

	// Start listening on port 8000
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 8000...")

	// Accept incoming connections in a loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Connection error:", err)
			continue
		}
		go rpc.ServeConn(conn) // Handle the connection in a new goroutine
	}
}

// Call server_session() in main function to start the server
func main() {
	server_session()
}
