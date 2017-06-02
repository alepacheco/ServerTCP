# ServerTCP
Easy to use tcp server for golang using dependency injection


## Sample usage
```go
package main

import (
  "os"
  "bufio"
  "net"
  "fmt"
  "github.com/alepacheco/ServerTCP"
)

func main() {
  // Create a server instance
  server := ServerTCP.NewServer()
  
  // Start listening for new connections in port 6666 and pass a listening function
  go server.Listen(":6666", listenFunc)
  
  // Read for clients messafes and passes a read function
  go server.Read(readFun)
	
  // Broadcast a message to all clients
  server.Broadcast("Hello clients")

}

// readFun is executed when a new message is recived
func readFun(message ServerTCP.ClientMessage) {
	fmt.Println("", message.Data, "'")
  // You can reply using message.Client.Writer (*bufio.Writer)
  
}

// listenFunc will be executed when a new client is connected 
func listenFunc(conn net.Conn) {
	fmt.Println("New connection: ", conn.RemoteAddr().String())
}
```
