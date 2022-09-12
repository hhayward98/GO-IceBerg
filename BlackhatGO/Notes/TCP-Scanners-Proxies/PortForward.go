// Proxying a TCP Client

// running this code on a proxy server

package main

import (
	"fmt"
	"log"
	"net"
)


func Handler(src net.Conn) {
	// destination forwarding to net.Dial("tcp", "Destination.com:80")
	destination, err := net.Dial("tcp", "google.com:80")
	if err != nil{
		log.Fatalln("Unable to connect to host")
	}

	defer destination.Close()

	// Run in goroutine to prevent io.Copy from blocking
	go func() {
		if _, err := io.Copy(destination, src); err != nil {
			log.Fatalln(err)
		}
	}()
	// Copy destination output back to out src

	if _, err := io.Copy(src, destination); err != nil {
		log.Fatalln(err)
	}
}


func main() {

	// Listen to port 80
	Listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln("Unable to bind Port")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalln("unable to accept connection")
		}
		go Handle(conn)
	}

}


