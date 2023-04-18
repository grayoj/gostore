# Distributed Key-Value Store

This is a simple implementation of a distributed key-value store in Golang. It uses TCP sockets for communication between clients and a server.

## Setup

To use the key-value store, follow these steps:

1. Clone the repository: git clone https://github.com/example/distributed-keyvalue-golang.git
2. Navigate to the project directory: cd distributed-keyvalue-golang
3. Build the project. Run `go build`
4. Start the server: ./distributed-keyvalue-golang
5. Start up client: ./distributed-keyvalue-golang

*Notice*: The server will listen on a _random TCP port_. The client will connect to the server on localhost and the port that the server is listening on.

## Using the Key-Value Store

1. Start the server.
```bash
./distributed-keyvalue-golang
```
2. Send a set message to the server to set a key-value pair.
```bash
./distributed-keyvalue-golang set key1 value1
```
3. To get the value of a key, send a `get` message.
```bash
./distributed-keyvalue-golang get key1
```

## Functionality

This program (code) makes use of different types to facilitate its functionality. The KeyValue type is used to represent key-value pairs, while the ServerState type is used to store the server's internal state and ensure synchronization using a mutex. Additionally, the Server type is used to manage the server itself and includes several methods, such as handling client connections and processing "get" and "set" messages.

The NewServer function creates a new Server instance, sets it to listen for TCP connections on a random port, initializes the server's encoder and decoder, and then returns the created server. On the other hand, the NewClient function creates a new client, connects to the server using the provided address, and then returns an encoder and decoder for the connection.

In the main function, a new Server is created and started in a new goroutine. The program waits for the server to start listening for connections before creating a new client. The client then sends a "set" message to the server to store a key-value pair, followed by a "get" message to retrieve thevalue for that key. Finally, the program prints the server's response and stops the server.

## Contributions

Feel free to open a PR or make an issue!

## References

https://github.com/akritibhat/Distributed-Key-Value-Store-Go-Lang
https://mrkaran.dev/posts/barreldb/
https://github.com/geohot/minikeyvalue
