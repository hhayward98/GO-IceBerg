/*Hello World*/

/*Registering a request Handler*/

/* Create a Handler that receives all incoming HTTP connections from browsers, HTTP clients or API request*/

/*http.ResponseWriter is where you write your text/html response to*/
/*http.Request contains all info about this HTTP request including the URL and header fields*/

// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello, URL Path: %s\n", r.URL.Path)
// })


/*Listen for HTTP Connections*/
// http.ListenAndServe(":80", nil)

/*Main*/

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}

