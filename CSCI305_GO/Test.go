package main


import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"bufio"
	"strings"
	"unicode"
)

// reading in contents from file

func ReadFile(FileName string) string {

	filename := FileName

	filebuffer, err := ioutil.ReadFile(filename)
	Debugger(err)

	inputdata := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(inputdata))
	data.Split(bufio.ScanRunes)
	var file_contents string
	file_contents = ""
	for data.Scan() {
		file_contents = file_contents + data.Text()

	}
	return file_contents
}




// This is the function for parsing out the numbers
func No_Numbers(data string) string {
	var non string
	non = ""
	for _, char := range data {
		if (!unicode.IsDigit(char)) {
			non = non + string(char)
		}
	}
	return non

}


// the function for writing the parsed string to a file
func writeTofile(contents string, fileoutname string) {

	file, err := os.Create(fileoutname)

	
	if err != nil {
		log.Fatalf("Error! task failed at creating file: %s", err)
	}

	defer file.Close()

	_ , err = file.WriteString(contents)

	Debugger(err)
}


func Debugger(err error) {
	if err != nil {
		panic(err)
	}
}


func main() {

	args := os.Args[1:]
	InputFile := args[0]
	fmt.Print("Enter the output filename : ")
	var outputfile string
	fmt.Scanln(&outputfile)

	var H string
	H = ReadFile(InputFile)
	H = No_Numbers(H)
	writeTofile(H, outputfile)
}