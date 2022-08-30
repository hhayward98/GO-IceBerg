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

func InitTemplates(Num int) {
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