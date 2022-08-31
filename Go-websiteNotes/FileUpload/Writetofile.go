package main 

import (
	"bufio"
	// "reflect"
	"strings"
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

func LineByLine() {

	var StrBuff string

	f, err := os.Open("test.go")
	Pancheck(err)

	defer f.Close()
	

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		Temp := scanner.Text()

		if strings.Contains(Temp, "main()") {

			fmt.Println(Temp)
		}

		// fmt.Println(reflect.TypeOf(Temp))  
	}


	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func CreateFile() {
	f, err := os.Create("./test.go")
	_, err2 := f.WriteString("package main\n\nimport ()\n\nfunc Home(w http.ResponseWriter, r *http.Response) {\n\tfmt.Println(`Home`)\n}\n\nfunc main() {\n\n\n\tfmt.Println(`hello`)\n}\n\n")
	Pancheck(err)
	Pancheck(err2)
	defer f.Close()

	fmt.Println("Done")
}

func CreateDir() {
	if err := os.Mkdir("static", os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func InitAppHead() {
	fmt.Println("Creating webapp File....")

	f, err := os.Create("./app.go")
	_, err2 := f.WriteString("package main\n\nimport ()\n\nfunc Home(w http.ResponseWriter, r *http.Response) {\n\tfmt.Println(`Home`)\n}\n\nfunc main() {\n\tfmt.Println(`hello`)\n}\n\n")
	Pancheck(err)
	Pancheck(err2)

	defer f.Close()
	fmt.Println("Done")

}

func InitDocker(){

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

func WriteCSS() {
	fmt.Println("Creating CSS....")
	if err := os.Mkdir("./static/css/", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	f, err := os.Create("./static/css/main.css")
	_, err2 := f.WriteString("html {\n\t\n\t\n}\n\nhead {\n\t\n\t\n}\n\nbody {\n\t\n\t\n}\n\nfooter {\n\t\n\t\n}\n")
	Pancheck(err)
	Pancheck(err2)

	defer f.Close()

}

func Options(){
	fmt.Println("Enter 1 for HTML ")
	fmt.Println("Enter 2 for CSS ")
	fmt.Println("Enter 4 for App.go")
	fmt.Println("Enter 3 for Dockerfile")

}

func main() {

	var CHOICE int
	var TempNum int 
	fmt.Println("Starting....")

	fmt.Println("Welcome!!")
	fmt.Println("Select option")
	Options()
	fmt.Scanln(&CHOICE)
	


	if CHOICE == 1 {
		fmt.Println("Enter the number of pages you want")
		fmt.Scanln(&TempNum)
		InitTemplates(TempNum)
	}else if CHOICE == 2 {
		WriteCSS()
	}else if CHOICE == 3 {
		InitAppHead()
	}else if CHOICE == 4 {
		InitDocker()
	}else if CHOICE == 5 {
		CreateFile()
		LineByLine()
	}

}

//Create unit test that test everything 
// Create limited functionality and unit tests