package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

type KeyValue struct {
	Key   string
	Value string
}

type ServerState struct {
	kvStore map[string]string // Internal key-value store
	mutex   sync.RWMutex      // Mutex for synchronization
}

// Define a type for the server
type Server struct {
	state    *ServerState   // Server state
	listener net.Listener   // Network listener
	encoder  *gob.Encoder   // Gob encoder for encoding messages
	decoder  *gob.Decoder   // Gob decoder for decoding messages
	stopChan chan bool      // Channel for stopping the server
	wg       sync.WaitGroup // Wait group for handling connections
}

// Start the server
func (s *Server) Start() {
	for {
		// Accept a connection from a client
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		s.wg.Add(1)
		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		s.wg.Done()
	}()

	decoder := gob.NewDecoder(conn)
	encoder := gob.NewEncoder(conn)

	var message KeyValue
	err := decoder.Decode(&message)
	if err != nil {
		log.Println(err)
		return
	}

	switch message.Key {
	case "get":
		s.handleGet(message, encoder)
	case "set":
		s.handleSet(message)
	default:
		log.Printf("unknown key: %s", message.Key)
	}
}

func (s *Server) handleGet(message KeyValue, encoder *gob.Encoder) {
	// Acquire a read lock on the mutex
	s.state.mutex.RLock()
	defer s.state.mutex.RUnlock()

	// Get the value for the key from the internal key-value store
	value, ok := s.state.kvStore[message.Value]

	// Encode the response message
	response := KeyValue{Value: value, Key: "get"}
	if ok {
		response.Key = "ok"
	}
	err := encoder.Encode(response)
	if err != nil {
		log.Println(err)
	}
}

// Handle a "set" message from a client
func (s *Server) handleSet(message KeyValue) {
	// Acquire a write lock on the mutex
	s.state.mutex.Lock()
	defer s.state.mutex.Unlock()

	// Set the value for the key in the internal key-value store
	s.state.kvStore[message.Key] = message.Value
}

func (s *Server) Stop() {
	s.listener.Close()
	close(s.stopChan)
	s.wg.Wait()
}

func NewServer() (*Server, error) {
	// Listen for TCP connections on a random port
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, err
	}

	// Print the port number for the listener
	fmt.Printf("Listening on port %d...\n", listener.Addr().(*net.TCPAddr).Port)

	state := &ServerState{
		kvStore: make(map[string]string),
		mutex:   sync.RWMutex{},
	}

	server := &Server{
		state:    state,
		listener: listener,
		encoder:  nil,
		decoder:  nil,
		stopChan: make(chan bool),
	}

	// Initialize the encoder and decoder for the server
	server.encoder = gob.NewEncoder(nil)
	server.decoder = gob.NewDecoder(nil)

	return server, nil

}

// Create a new client
func NewClient(address string) (*gob.Encoder, *gob.Decoder, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, nil, err
	}

	encoder := gob.NewEncoder(conn)
	decoder := gob.NewDecoder(conn)

	return encoder, decoder, nil

}

func main() {
	server, err := NewServer()
	if err != nil {
		log.Fatal(err)
	}
	go server.Start()

	<-time.After(1 * time.Second)

	encoder, decoder, err := NewClient(fmt.Sprintf("localhost:%d", server.listener.Addr().(*net.TCPAddr).Port))
	if err != nil {
		log.Fatal(err)
	}

	// Set a key-value pair
	message := KeyValue{Key: "set", Value: "key1=value1"}
	err = encoder.Encode(message)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve the value for the key
	message = KeyValue{Key: "get", Value: "key1"}
	err = encoder.Encode(message)
	if err != nil {
		log.Fatal(err)
	}

	var response KeyValue
	err = decoder.Decode(&response)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("response: %s\n", response.Value)

	server.Stop()
}

