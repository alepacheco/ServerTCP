package ServerTCP

import "net"

type FunctionListen func(net.Conn)
type FunctionRead func(ClientMessage)

type Server struct {
	clients []*Client
	channel chan *ClientMessage
}

// Broadcast sends a string to all clints outgoing chain
func (server *Server) Broadcast(data string) {
	for _, client := range server.clients {
		client.Write(data)
	}
}

func (server *Server) Join(connection net.Conn) {
	client := NewClient(connection, server)
	server.clients = append(server.clients, client)
	//go func() { for { server.incoming <- <-client.incoming } }()
}

func (server *Server) Listen(port string, fun FunctionListen) {
	listener, _ := net.Listen("tcp", port)
	for {
		// When new connections add them to the list
		conn, _ := listener.Accept()
		server.Join(conn)
		fun(conn)
	}
}

func (server *Server) Read(fun FunctionRead) {
	for {
		fun(*<-server.channel)
	}
}

func NewServer() *Server {
	server := &Server{
		clients: make([]*Client, 0),
		channel: make(chan *ClientMessage),
	}
	return server
}
