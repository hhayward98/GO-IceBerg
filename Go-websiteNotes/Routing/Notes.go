/*Using gorilla/mux package*/


/*Creating a new Router*/

/*Create a new request router(main router of web app)*/

/*Recives all HTTP connections and pass to request handlers*/

// RouterName := mux.NewRouter()

/*instead of calling http.HandleFunc(), use RouterName.HandleFunc() */

/*URL parameters allow for exctract dynamic segments*/

// RouterName.HandleFunc("/books/<parameter1>/page/<parameter2>", func(w http.ResponseWriter, r *http.Request){
// 	//get data using parameter1 
// 	// navigate data using parameter2
// })

/*Get the data using mux.Vars()*/

// func(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(RouterName)
// 	vars["parameter1"] // data
// 	vars["parameter2"] // segment of data
// }

/*Setting the HTTP servers router*/

// http.ListenAndServe(":80", RouterName)


/*Methods*/

/*Restrict the request handler to specific HTTP methods*/



/*
RouterName.HandleFunc("/books/<paramater1>", CreateBook).Methods("POST")
RouterName.HandleFunc("/books/<paramater1>", ReadBook).Methods("GET")
RouterName.HandleFunc("/books/<paramater1>", UpdateBook).Methods("PUT")
RouterName.HandleFunc("/books/<paramater1>", DeleteBook).Methods("DELETE")
	*/

/*Hostnames & Subdomains*/

/*
RouterName.HandleFunc("/books/<paramerter1", BookHandeler).Host("www.mybookstore.com")

*/


/*Schemes*/

/* Restrict the request handler to http/https*/
/*
RouterName.HandlerFunc("/secure", SecureHandler).Schemes("hettps")
RouterName.HandleFunc("/insecure", InsecureHandler).Schemes("http")
*/

/*Path Prefixes & Subrouters*/
/*Restrict the request handler to specific pathj prefixes*/
/*
BookRouter := RouterName.PathPrefix("/books").Subrouter()
Bookrouter.HandleFunc("/", AllBooks)
BookRouter.HandleFunc("/<paramater2", GetBook)
*/

/*Main*/

package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"

)

func main() {
	Rmain := mux.NewRouter()

	Rmain.HandleFunc("/collections/<CollectionName>/TokenID/<TokenID>", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(Rmain)
		CName := vars["CollectionName"]
		TokenID := vars["TokenID"]

		fmt.Fprintf(w, "Request for TokenID: %s from %s\n", TokenID, CName)
	})

	http.ListenAndServe(":80", Rmain)
}