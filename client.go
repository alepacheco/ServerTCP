package ServerTCP

import (
	"bufio"
	"net"
)

type Client struct {
	Reader    *bufio.Reader
	Writer   *bufio.Writer
	readChannel    chan *ClientMessage
}
type ClientMessage struct {
	Client	*Client
	Data	string
}

func (client *Client) Listen() {
	for {
		line, _ := client.Reader.ReadString('\n')
		
		message := &ClientMessage{
			Client: client,
			Data:	line,
		}
		client.readChannel <- message
	}
}

func (client *Client) Write(data string) {
	client.Writer.WriteString(data)
	client.Writer.Flush()
}

func NewClient(connection net.Conn, server *Server) *Client {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	client := &Client {
		Reader: reader,
		Writer: writer,
		readChannel: server.channel,
	}
	go client.Listen()

	return client
}