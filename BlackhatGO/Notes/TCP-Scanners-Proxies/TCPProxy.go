package main

import (
	"fmt"
	"log"
	"os"

)


// defines an io.Reader to read from stdin
type FReader struct{}


// reads data from stdin
func (reader *FReader) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	// 
	return os.Stdin.Read(b)

}

// FWriter defines an io.Write to Stdin
type FWriter struct{}

func (writer *FWriter) Write(b []byte) (int error) {
	fmt.Print("out> ")
	return os.Stdin.Write(b)
}

//  Go io package contains Copy function below
// func Copy(dst io.Writer, src io.Reader) (written int64, error)

func GOCOPY() {
	var (
		reader FReader
		writer FWriter
	)

	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}

}

func main() {

	var (
		reader FReader
		writer FWriter
	)

	// make buffer 
	input := make([]byte, 4096)


	// read input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("unable to read data")
	}

	fmt.Println("Read %d bytes from stdin\n", s)

	// Write output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("unable to Write data")
	}

	fmt.Printf("wrote %d bytes to stdout/n", s)

}