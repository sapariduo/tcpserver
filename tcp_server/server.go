package tcp_server

import (
	"bufio"
	"crypto/tls"
	"errors"
	"fmt"
	"log"
	"net"
)

// Client holds info about connection
type Client struct {
	conn   net.Conn
	Server *server
}

// TCP server
type server struct {
	address                  string // Address to open connection: localhost:9999
	config                   *tls.Config
	onNewClientCallback      func(c *Client)
	onClientConnectionClosed func(c *Client, err error)
	onNewMessage             func(c *Client, message string)
	onEmptyMessage           func(c *Client, err error)
}

// Read client data from channel
func (c *Client) listen() {
	c.Server.onNewClientCallback(c)
	reader := bufio.NewReaderSize(c.conn, 1024)

	// reader := bufio.NewReader(c.conn)
	for {
		prefix, err := reader.Peek(2)
		if err != nil {
			c.conn.Close()
			c.Server.onClientConnectionClosed(c, err)
			return
		}
		if len(prefix) == 0 || fmt.Sprintf("%x", prefix) == "0a0a" {
			err := errors.New("error on empty data")
			c.conn.Close()
			c.Server.onEmptyMessage(c, err)
			return
		}
		message, _ := reader.ReadString(0)
		c.Server.onNewMessage(c, message)
	}
}

// Send text message to client
func (c *Client) Send(message []byte) error {
	_, err := c.conn.Write(message)
	return err
}

// Send bytes to client
func (c *Client) SendBytes(b []byte) error {
	_, err := c.conn.Write(b)
	return err
}

func (c *Client) Conn() net.Conn {
	return c.conn
}

func (c *Client) Close() error {
	return c.conn.Close()
}

// Called right after server starts listening new client
func (s *server) OnNewClient(callback func(c *Client)) {
	s.onNewClientCallback = callback
}

// Called right after connection closed
func (s *server) OnClientConnectionClosed(callback func(c *Client, err error)) {
	s.onClientConnectionClosed = callback
}

// Called when Client receives new message
func (s *server) OnNewMessage(callback func(c *Client, message string)) {
	s.onNewMessage = callback
}

func (s *server) OnEmptyMessage(callback func(c *Client, err error)) {
	s.onEmptyMessage = callback
}

// Listen starts network server
func (s *server) Listen() {
	var listener net.Listener
	var err error
	if s.config == nil {
		listener, err = net.Listen("tcp4", s.address)
	} else {
		listener, err = tls.Listen("tcp4", s.address, s.config)
	}
	if err != nil {
		log.Fatal("Error starting TCP server.")
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		client := &Client{
			conn:   conn,
			Server: s,
		}
		go client.listen()
	}
}

// Creates new tcp server instance
func New(address string) *server {
	log.Println("Creating server with address", address)
	server := &server{
		address: address,
		config:  nil,
	}

	server.OnNewClient(func(c *Client) {})
	server.OnNewMessage(func(c *Client, message string) {})
	server.OnClientConnectionClosed(func(c *Client, err error) {})

	return server
}

func NewWithTLS(address string, certFile string, keyFile string) *server {
	log.Println("Creating server with address", address)
	cert, _ := tls.LoadX509KeyPair(certFile, keyFile)
	config := tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	server := &server{
		address: address,
		config:  &config,
	}

	server.OnNewClient(func(c *Client) {})
	server.OnNewMessage(func(c *Client, message string) {})
	server.OnClientConnectionClosed(func(c *Client, err error) {})

	return server
}
