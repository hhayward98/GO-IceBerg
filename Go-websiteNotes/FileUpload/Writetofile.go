package main 

import (
	// "bufio"
	"fmt"
	"os"
	"log"
	"strconv"

)

func Pancheck(e error) {
	if e != nil {
		panic(e)
	}
}

func CreateFile() {
	f, err := os.Create("./test.go")
	Pancheck(err)

	defer f.Close()

	fmt.Println("Done")
}

func CreateDir() {
	if err := os.Mkdir("static", os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func InitDocker(){

	// ask user for File name and replace AEsir with that var
	fmt.Println("Creating Docker File....")
	f, err := os.Create("./Dockerfile")
	_, err2 := f.WriteString("FROM golang:1.18\n\nRUN mkdir /GoWeb\n\nADD . /GoWeb\n\nWORKDIR /GoWeb\n\nCOPY go.* ./\n\nRUN go mod download && go mod verify\n\nRUN go build -o app .\n\nEXPOSE 8080\n\nCMD ['/GoWeb/app']")
	Pancheck(err)
	Pancheck(err2)

	defer f.Close()
	fmt.Println("Done")
}

func InitTemplates(Num int) {
	fmt.Println("Creating Templates....")
	if err := os.MkdirAll("./static/templates/", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if Num == 1 {
		WriteHTML("index.html")
	}else if Num > 1 {
		WriteHTML("index.html")

		for i := 0; i < Num; i++ {
			strBuf := "page"
			s1 := strconv.Itoa(i)
			strBuf += s1
			strBuf += ".html"
			fmt.Println(strBuf)
			WriteHTML(strBuf)
		}
	}

	fmt.Println("Done")
}

func WriteHTML(Fname string) {

	f, err := os.Create("./static/templates/"+ Fname)
	_, err2 := f.WriteString("<!DOCTYPE html>\n<html>\n<head>\n\t<meta charset='utf-8'>\n\t<meta name='viewport' content='width=device-width, initial-scale=1'>\n\t<meta http-equiv='X-UA-Compatible' content='ie=edge' />\n\t<title>Home</title>\n</head>\n<body>\n\t<h3>Title</h3>\n\t<p>Information</p>\n</body>\n</html>")
	Pancheck(err)
	Pancheck(err2)

	defer f.Close()

}

func main() {

	fmt.Println("Starting....")
	// InitTemplates(3)

}