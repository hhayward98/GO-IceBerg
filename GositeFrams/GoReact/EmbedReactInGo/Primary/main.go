package main

import (

	"log"
	"embed"
	"io/fs"
	"net/http"
	
)


var embeddedFiles embed.FS


func getFileSystem() http.FileSystem {


	Fsys, err := fs.Sub(embeddedFiles, "ui/build")
	if err != nil {
		log.Fatal(err);
	}
	log.Println(Fsys)

	data, _ := embeddedFiles.ReadFile("./Home.html")

	log.Println(data)

	return http.FS(Fsys)


}


func main() {


	log.Println("Listening....")
	http.Handle("/", http.FileServer(getFileSystem()))
	log.Fatal(http.ListenAndServe(":8080", nil))
}