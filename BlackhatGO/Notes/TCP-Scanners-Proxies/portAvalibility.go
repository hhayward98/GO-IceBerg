package main

import (
	"fmt"
	"net"
)


func nonconcurrent() {

	for i := 1; i <= 1024; i {
		addy := fmt.Sprintf("scanme.nmap.org:%d", i)
		fmt.Println(addy)
	}

}

func test1() {

	_, err := net.Dial("tcp","scanme.nmap.org:80")

	if err == nil {
		fmt.Println("Connection Success!!")
	}else {
		fmt.Println(err)
	}
}

func test2() {
	for i := 1; i <= 1024; i++ {
		addy := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", addy)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}

// Concurrent Scanning
func test3() {
	//too fast
	for i := 1; i <= 1024; i++ {
		go func(j int) {
			addy := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
}

func ConcurrentWait() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			addy := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", addy)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
}


func main() {
	fmt.Println("Starting....")


}