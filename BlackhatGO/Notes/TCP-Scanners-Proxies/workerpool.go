package main

import (
	"fmt"
	"sync"

)


func Worker(ports chan int, wg *sync.WaitGroup) {

	for PT := range ports {
		fmt.Println(PT)
		wg.Done()
	}
}


func main() {
	//Create channel to be buffered
	ports := make(chan int, 100)

	var wg sync.WaitGroup
	// for loop to create desiered number of workers
	for i := 0; i < cap(ports); i++ {
		go Worker(ports, &wg)
	}
	// sending ports on ports channel to the worker
	for i := 1; i <= 1024; i++ {
		wg.Add()
		ports <- i
	}
	wg.Wait()
	close(ports)

}