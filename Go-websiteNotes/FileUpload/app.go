package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

func Debugger(err error) {
	if err != nil{
		log.Fatal(err)
	}
}

func UploadFile(w http.ResponseWriter, r *http.Request) {

	// parse multipart form.
	// 10 << 20 specifies max upoload of 10 MB files
	r.ParseMultipartForm(10 << 20)

	// FormFile returns first file for given key
	// also returns file header for the filename, header, and size of file.
	file, handler, err := r.FormFile("myfile")
	if err != nil {
		fmt.Println("Error getting file")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Println("Uploaded File: %+v\n", handler.Filename)
	fmt.Println("File Size: %+v\n", handler.Size)
	fmt.Println("MIME Header: %+v\n", handler.Header)


	TempFile, err := ioutil.TempFile("Temp-folder", "upload-*.png")
	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()

	//read all contents from uploadfile into a byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	//write byte array to tempfile
	tempFile.Write(fileBytes)

	fmt.Println(w, "Successfully Uploaded\n")

}

func AppRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}

func main() {
	AppRoutes()
}

