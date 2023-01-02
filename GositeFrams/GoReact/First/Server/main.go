package main 

import (
	"log"
	"net/http"

	"github.com/rs/cors"

)


func Route1(w http.ResponseWriter, r *http.Request) {
	log.Println("Route one Request made....")

	String := "Hello From RouteOne!!"
	w.Header().Set("Content-Type", "application/json")
	

	log.Println("Method: ", r.Method)

	// log.Println("ResponseWriter \n",w)
	// log.Println("\nRequest \n", r)



	w.Write([]byte("{\"Message\": \""+String+"\"}"))

}


func Route2(w http.ResponseWriter, r *http.Request) {
	log.Println("Route two Request Made....")

	// Cors 
	w.Header().Set("Content-Type", "application/json")

	String := "Hello From RouteTwo!!"
	if (r.Method == "POST") {
		log.Println("Post method used")
		w.Write([]byte("{\"Message\": \""+String+"\", \"Auth\": \"true\"}"))
		return
	} else if (r.Method == "GET") {
		log.Println("Get Method Used")
		w.Write([]byte("{\"Message\": \""+String+"\", \"Auth\": \"false\"}"))
		return
	} else {
		log.Println("Method Used: ", r.Method)
		w.Write([]byte("{\"Message\": \""+String+"\", \"Auth\": \"false\"}"))
		return
	}


}

func Route3(w http.ResponseWriter, r *http.Request) {
	log.Println("Route three Request Made....")

	w.Header().Set("Content-Type", "application/json")


	log.Println("\nRequest \n", r)


}


func Route4(w http.ResponseWriter, r *http.Request) {
	log.Println("Route four Request Made....")

	w.Header().Set("Content-Type", "application/json")

	// Only GET Method Allowed
	if r.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}



}


func RouteP(w http.ResponseWriter, r *http.Request) {
	log.Println("Route Post Request Made....")

	w.Header().Set("Content-Type", "application/json")

	if (r.Method == "POST") {

		w.Write([]byte("{\"Message\": \"Post Method Only!\", \"Auth\": \"true\"}"))
		return
	} else {
		log.Println("Method: ", r.Method)
		w.Write([]byte("{\"Message\": \"Method is Not Valid\", \"Auth\": \"false\"}"))
	}

}



func AppRoutes() {

	mux := http.NewServeMux()

	C := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{http.MethodPost,http.MethodGet},
		Debug: true,
	})


	mux.HandleFunc("/RouteOne", Route1)
	mux.HandleFunc("/RouteTwo", Route2)
	mux.HandleFunc("/RouteThree", Route3)
	mux.HandleFunc("/RouteFour", Route4)
	mux.HandleFunc("/RouteP", RouteP)



	handler := C.Handler(mux)

	log.Fatal(http.ListenAndServe(":8080", handler))



}


func main() {

	log.Println("Listening.....")

	AppRoutes()


}