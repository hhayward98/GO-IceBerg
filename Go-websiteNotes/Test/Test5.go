package main

import (
	"bufio"
	"fmt"
	"strings"
	"os"
	"log"
	"strconv"
)

func Debugger(err error) {
	if err != nil {
		log.Fatal(err)
	}
}


func AppendToMain(Num int) {
	// append Temp(each line) to the string buffer
	// can make conditions to search for content
	// if Temp contains main() append to StrBuff and stepinto condition
		// after appending current line from file, we can inject code to run in main()


	var strbuffer string
	var status bool = false
	f, err := os.Open("Test.go")
	Debugger(err)

	scanner := bufio.NewScanner(f)



	for scanner.Scan() {

		Line := scanner.Text()
		strbuffer += Line +"\n"

		if strings.Contains(Line, "main()"){
			if status == false {
				fmt.Println(Line)
				StrInj := "\n\n\thttp.HandleFunc('/', Home)"
				FInject := strings.ReplaceAll(StrInj, "'", `"`,)
				strbuffer += FInject + "\n"
				status = true
			}

			if Num != 1 {

				for j := 0; j < Num; j++ {
					StJ := strconv.Itoa(j)
					strinj := "\thttp.HandleFunc('/', Page"+StJ+")\n"
					Rinject := strings.ReplaceAll(strinj, "'", `"`,)
					//append to string Buffer code for injection
					strbuffer += Rinject 

				}

			}

		}

	}

	h, err := os.Create("./Test.go")
	_, err2 := h.WriteString(strbuffer)
	Debugger(err)
	Debugger(err2)

	defer f.Close()
	defer h.Close()


	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Done")
}

func main() {

	AppendToMain(2)

	Tem := "FROM golang:1.18\n\nRUN mkdir /GoWeb\n\nADD . /GoWeb\n\nWORKDIR /GoWeb\n\nCOPY go.* ./\n\nRUN go mod download && go mod verify\n\nRUN go build -o app .\n\nEXPOSE 8080\n\nCMD ['/GoWeb/app']"
	New := strings.ReplaceAll(Tem, "'", `"`,)
	f, err := os.Create("./Dockerfile")
	_, err2 := f.WriteString(New)
	Debugger(err)
	Debugger(err2)

	defer f.Close()
	fmt.Println("Done")

}