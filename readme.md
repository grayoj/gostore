# Distributed Key-Value Store

This is a simple implementation of a distributed key-value store in Golang. It uses TCP sockets for communication between clients and a server.

## Setup

To use the key-value store, follow these steps:

1. Clone the repository
2. Navigate to the project directory: cd gostore
3. Build the project. Run `go build`
4. Start the server: ./gostore
5. Start up client: ./gostore

*Notice*: The server will listen on a _random TCP port_. The client will connect to the server on localhost and the port that the server is listening on.

## Using the Key-Value Store

Enable permissions on your machine. `chmod +x gostore``

1. Start the server.
```bash
./gostore
```
2. Send a set message to the server to set a key-value pair.
```bash
./gostore set key1 value1
```
3. To get the value of a key, send a `get` message.
```bash
./gostore get key1
```

## Functionality

This program (code) makes use of different types to facilitate its functionality. The `KeyValue` type is used to represent key-value pairs, while the ServerState type is used to store the server's internal state and ensure synchronization using a mutex. Additionally, the Server type is used to manage the server itself and includes several methods, such as handling client connections and processing "get" and "set" messages.

The `NewServer` function creates a new Server instance, sets it to listen for TCP connections on a random port, initializes the server's encoder and decoder, and then returns the created server. On the other hand, the `NewClient` function creates a new client, connects to the server using the provided address, and then returns an encoder and decoder for the connection.

In the main function, a new Server is created and started in a new goroutine. The program waits for the server to start listening for connections before creating a new client. The client then sends a "set" message to the server to store a key-value pair, followed by a "get" message to retrieve thevalue for that key. Finally, the program prints the server's response and stops the server.

## Contributions

Feel free to open a PR or make an issue!

## References

1. https://github.com/akritibhat/Distributed-Key-Value-Store-Go-Lang
2. https://mrkaran.dev/posts/barreldb/
3. https://github.com/geohot/minikeyvalue

## License

MIT License

Copyright (c) 2023 Gerald Maduabuchi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
