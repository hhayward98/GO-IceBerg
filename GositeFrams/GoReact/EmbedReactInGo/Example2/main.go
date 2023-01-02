package main

import (
	"log"
	"io"
	"os"
	"fmt"
	"mime"
	"path/filepath"
	"strings"
	"net/http"
	
	// "io/fs"
	// app "github.com/somthingTODO"
)


var uiFS fs.FS

func init() {

	var err error
	uiFS, err = fs.Sub(app.UI, "ui/build")
	if err != nil {
		log.Fatal("Failed To get UI FS: ", err)
	}
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {

	// TODO:
}

func handleApi(w http.ResponseWriter, r *http.Request) {

	// TODO:
}

func handleStatic(w http.ResponseWriter, r *http.Request) {

	// Only GET Method Allowed

	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}


	Fpath := filepath.Clean(r.URL.Path)
	if Fpath == "/" {	
		//  Include all paths that are routed on the UI side
		Fpath = "index.html"
	}

	Fpath = strings.TrimPrefix(Fpath, "/")

	file, err := uiFS.Open(Fpath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("file", Fpath, " not found:", err)
			http.NotFound(w,r)
			return
		}

		log.Println("file", Fpath, "Cannot be read:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	contentType := mime.TypeByExtension(filepath.Ext(Fpath))
	w.Header().Set("Content-Type", contentType)
	if strings.HasPrefix(Fpath, "statix/") {
		w.Header().Set("Cache-Control", "public, max-age=31536000")
	}
	stat, err := file.Stat()
	if err == nil && stat.Size() > 0 {
		w.Header().Set("Conent-Length", fmt.Sprintf("%d", stat.Size()))
	}

	n, _ := io.Copy(w,file)
	log.Println("file", Fpath, "copied", n,"bytes")



}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/HealthCheck", handleHealthCheck)

	// API calls
	mux.HandleFunc("/api", handleApi)

	// Static Page
	mux.HandleFunc("/", handleStatic)


	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Println("server failed:", err)
	}

}