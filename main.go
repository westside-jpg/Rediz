package main

import (
	"fmt"
	"io"
	"log"
	"net"
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

		buf := make([]byte, 1024)
		answer := []byte("+PONG\r\n")

		go func() {
			for {
				n, err := conn.Read(buf)
				if err != nil {
					if err != io.EOF {
						log.Printf("Listener reading error on port %v: %v", lst.Addr(), err)
					}
					break
				}

				fmt.Printf("Client send: %v\n", string(buf[:n]))

				_, err = conn.Write(answer)
				if err != nil {
					fmt.Printf("Error writing answer: %v", err)
					break
				}
			}
		}()

	}
}
