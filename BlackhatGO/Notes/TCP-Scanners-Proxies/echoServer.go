package main

import (
	"fmt"
	"net"
)



// echo data

func echo(conn net.Conn) {
	defer conn.Close()
	

	// Create buffer for data 
	buffer := make([]byte, 512)
	// use buffer to read and write data from and to the connection
	for {
		// receive data with conn.Read into buffer
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client Disconnected")
			break
		}
		if err != nil {
			log.Println("unexpected Error")
			break
		}
		log.Printf("%d bytes recived: %s\n", size, string(b))

		// send data with conn.Write
		log.Println("writing data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("unable to Write data")
		}
	}
}

func main() {
	// Bind to TCP port on all interfaces
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to bind to port")
	}
	log.Println("Listening on 0.0.0.0:20080")

	// an infinite loop keeps the server listening for connections 
	for {
		// wait for connection. create net.conn
		conn, err := listener.Accept()
		log.Println("Received Connection")
		if err != nil {
			log.Fatalln("unable to accept connection")
		}

		// handling connection using goroutine for concurrency
		go echo(conn)
	}

}




// using a Buffer Listener

func echo2(conn net.Conn) {
	defer conn.Close()

	// init new buffered Reader and Writer via NewReader(io.Reader) and NewWriter(io.Writer)
	reader : = bufio.NewReader(con)
	s, err := reader.ReadString('\n')
	if err != nill {
		log.Fatalln("unable to read data")
	}

	log.Printf("Read %d bytes: %s", len(s), s)

	log.Println("Writing data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriteString(s); err != nil {
		log.Fatalln("Unable to write data")
	}

	// flush to write all the data to the underlying writer
	writer.Flush()
}


// using copy 

func echoCopy(conn net.Conn) {

	defer conn.Close()
	// use Copy method
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatalln("unable to read/Write data")
	}

}




