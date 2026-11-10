package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/westside-jpg/Rediz/parser"
	"github.com/westside-jpg/Rediz/types"
)

func main() {
	lst, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("Listener initialization error on port: %v", err)
	}
	defer func(lst net.Listener) {
		err := lst.Close()
		if err != nil {
			panic(err)
		}
	}(lst)

	fmt.Printf("Listening on port 1234\n")

	for {
		conn, err := lst.Accept()
		if err != nil {
			log.Fatalf("Listener error on port %v: %v", lst.Addr(), err)
		}
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				log.Printf("Connection close error on port %v: %v", lst.Addr(), err)
			}
		}(conn)

		answer := []byte("+PONG\r\n")
		go func() {
			request := types.Request{}
			for {
				reader := bufio.NewReader(conn)
				request, err = parser.RequestParser(reader)
				if err != nil && err == io.EOF {
					fmt.Print(err)
				}
				fmt.Printf("Request received: %v\n", request)

				_, err = conn.Write(answer)
				if err != nil {
					fmt.Printf("Error writing answer: %v", err)
					break
				}
			}
		}()

	}
}
