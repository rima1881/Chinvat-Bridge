package communication

import (
	"botzilla/pkg/command"
	"botzilla/pkg/message"
	"fmt"
	"net"
	"os"
)

func StartTCPServer() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("There was an error starting the server: \n", err)
		os.Exit(1)
	}

	defer listener.Close()

	fmt.Println("Botzilla has started on port 8080")

	// Router decides what to do with incomming commands
	router := command.NewCommandRouter()

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: \n", err)
			continue
		}

		go handleConnection(conn, router)
	}
}

func handleConnection(conn net.Conn, router *command.CommandRouter) {

	// Might need to change :O
	rawMessage := make([]byte, 1024)
	_, err := conn.Read(rawMessage)
	if err != nil {
		fmt.Println("Error reading command:\n", err)
		conn.Close()
		return
	}

	message, err := message.Deserialize(rawMessage)
	if err != nil {
		fmt.Println("There was an error decoding :\n", err)
	}

	router.Route(&message)

}
