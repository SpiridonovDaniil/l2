package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	r, _ := net.Listen("tcp", "localhost:8080")
	conn, _ := r.Accept()

	handler(conn)

	err := r.Close()
	if err != nil {
		log.Println(err)
	}
}

func handler(conn net.Conn) {
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		switch message {
		case "Ctrl+D":
			return
		default:
			fmt.Print("Message Received:", message)
			newMessage := strings.ToUpper(message)
			_, err := conn.Write([]byte(time.Now().String() + " " + newMessage + "\n"))
			if err != nil {
				log.Print(err)
			}
		}
	}
}
