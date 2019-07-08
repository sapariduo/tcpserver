package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/sapariduo/tcpserver/tcp_server"
	"github.com/sapariduo/tcpserver/utils"
)

func main() {
	host := flag.String("host", "0.0.0.0", "Host IP address")
	port := flag.String("port", "12345", "Port")
	flag.Parse()
	serverAddr := fmt.Sprintf("%s:%s", *host, *port)
	server := tcp_server.New(serverAddr)

	server.OnNewClient(func(c *tcp_server.Client) {
		// new client connected
		// lets log device
		log.Println("New_Connection", c.Conn().RemoteAddr())
	})
	server.OnNewMessage(func(c *tcp_server.Client, message string) {
		log.Printf("Message Length: %d", utils.HexT(message))
		msg := msgParser(string(message))

		switch msg.Type() {
		case 1:
			c.Send([]byte("0002640113BC"))
		case 68:
			c.Send([]byte("0002640113BC"))
		case 15:
			c.Send([]byte("00027301CB25"))
		}
	})
	server.OnClientConnectionClosed(func(c *tcp_server.Client, err error) {
		log.Println("Quit", c.Conn().RemoteAddr())
	})

	server.OnEmptyMessage(func(c *tcp_server.Client, err error) {
		log.Println(err, c.Conn().RemoteAddr())
	})

	server.Listen()
}
