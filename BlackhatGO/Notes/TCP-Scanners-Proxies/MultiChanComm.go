package main

import (
	"fmt"
	"net"
	"sort"
)


func Worker(ports, results chan int) {
	for PT := range ports {
		addy := fmt.Sprintf("scanme.nmap.org:%d", PT)
		conn, err := net.Dial("tcp", addy)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- PT
	}
}


func main() {
	ports := make(chan int, 100)
	results := make(chan int)
	var OPorts []int


	for i := 0; i < cap(ports); i++ {
		go Worker(ports, results)
	}

	go func() {
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()


	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			OPorts = append(OPorts, ports)
		}
	}

	close(ports)
	close(results)
	sort.Ints(OPorts)
	for _, port := range OPorts {
		fmt.Printf("%d open\n", port)
	}

}