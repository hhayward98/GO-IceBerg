package main 

import (
	"log"
	"net/http"

	"github.com/rs/cors"

)


func Route1(w http.ResponseWriter, r *http.Request) {

	log.Println("Route one Request made...")

	w.Header().Set("Content-Type", "application/json")
	

	log.Println("ResponseWriter \n",w)
	log.Println("\nRequest \n", r)

	log.Println("\n\n")
	log.Println("Method: ", r.Method)

	String := "Hello From RouteOne!!"


	w.Write([]byte("{\"Message\": \""+String+"\"}"))

}


func AppRoutes() {

	mux := http.NewServeMux()

	C := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{http.MethodPost,http.MethodGet},
		Debug: true,
	})


	mux.HandleFunc("/RouteOne", Route1)


	handler := C.Handler(mux)

	log.Fatal(http.ListenAndServe(":8080", handler))



}


func main() {

	log.Println("Listening.....")

	AppRoutes()


}