package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"log"
	"bufio"
	"strings"
	"unicode"
)

func main() {
	//here we are testing for file name by command line argument
	if len(os.Args) < 2 {
		fmt.Println("Error!, you forgot to set the file name in your argument.")
		return
	}

	// setting the file name
	filename := os.Args[1]

	// setting the output file name
	var outputfile string
	outputfile = "No_Numbers.txt"

	// reading in the data parsing it and writing the data back to a text file.
	var parseString string
	parseString = ReadFile(filename)
	parseString = No_Numbers(parseString)
	writeTofile(parseString, outputfile) 


}

// error cheacking function
func Debugger(err error) {
	if err != nil {
		panic(err)
	}
}


// reads in data from a file. 

func ReadFile(FileName string) string {

	filename := FileName

	filebuffer, err := ioutil.ReadFile(filename)
	// I had to use this debugger to check for err
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

// removes all the numbers from the text
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

// writes to a file
func writeTofile(contents string, fileoutname string) {

	file, err := os.Create(fileoutname)

	
	if err != nil {
		log.Fatalf("Error! task failed at creating file: %s", err)
	}

	defer file.Close()

	_ , err = file.WriteString(contents)

	Debugger(err)
}

































