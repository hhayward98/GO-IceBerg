package main

import (

	"log"
	"embed"
	"io/fs"
	"net/http"
)


// go :Embed Build
var embeddedFiles embed.FS


func getFileSystem() http.FileSystem {


	// get build subdirectory as the root directory
	// this allows it to be passed the http.FileServer

	Fsys, err := fs.Sub(embeddedFiles, "build")
	if err != nil {
		log.Fatal(err);
	}


	return http.FS(Fsys)


}


func main() {
	log.Println("Listening....")
	http.Handle("/", http.FileServer(getFileSystem()))
	log.Fatal(http.ListenAndServe(":8080", nil))
}