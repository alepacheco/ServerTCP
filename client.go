package ServerTCP

import (
	"bufio"
	"net"
)

type Client struct {
	reader    *bufio.Reader
	writer   *bufio.Writer
	readChannel    chan *ClientMessage
}
type ClientMessage struct {
	Client	*Client
	Data	string
}

func (client *Client) Listen() {
	for {
		line, _ := client.reader.ReadString('\n')
		line = line.strip()
		message := &ClientMessage{
			Client: client,
			Data:	line,
		}
		client.readChannel <- message
	}
}

func (client *Client) Write(data string) {
	client.writer.WriteString(data)
	client.writer.Flush()
}

func NewClient(connection net.Conn, server *Server) *Client {
	writer := bufio.NewWriter(connection)
	reader := bufio.NewReader(connection)

	client := &Client {
		reader: reader,
		writer: writer,
		readChannel: server.channel,
	}
	go client.Listen()

	return client
}