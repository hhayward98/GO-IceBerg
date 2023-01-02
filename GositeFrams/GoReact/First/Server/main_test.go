package main 

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	// "github.com/rs/cors"

)



func TestRoute1(t *testing.T) {

	res, err := http.Get("http://localhost:8080/RouteOne")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	ByteBuffer, err := ioutil.ReadAll(res.Body)
	got := string(ByteBuffer)
	log.Println("response data: ",got)

	want := `{"Message": "Hello From RouteOne!!"}`

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}

}

// TODO Finish Test

func TestRoute2G(t *testing.T) {

	res, err := http.Get("http://localhost:8080/RouteTwo")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	ByteBuffer, err := ioutil.ReadAll(res.Body)
	got := string(ByteBuffer)
	log.Println("response data: ",got)

	want := `{"Message": "Hello From RouteTwo!!", "Auth": "false"}`

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}

}

func TestRoute2P(t *testing.T) {

	res, err := http.Post("http://localhost:8080/RouteTwo", "application/json",  bytes.NewBuffer([]byte("{\"Message\": \"Test\"}")))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	ByteBuffer, err := ioutil.ReadAll(res.Body)
	got := string(ByteBuffer)
	log.Println("response data: ",got)

	want := `{"Message": "Hello From RouteTwo!!", "Auth": "true"}`

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}

}

func TestRoute3(t *testing.T) {

	res, err := http.Get("http://localhost:8080/RouteThree")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	ByteBuffer, err := ioutil.ReadAll(res.Body)
	got := string(ByteBuffer)
	log.Println("response data: ",got)

	want := "TBD"

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}

}

func TestRoute4(t *testing.T) {

	res, err := http.Get("http://localhost:8080/RouteFour")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	ByteBuffer, err := ioutil.ReadAll(res.Body)
	got := string(ByteBuffer)
	log.Println("response data: ",got)

	want := "TBD"

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}

}

func TestRoutePP(t *testing.T) {

	res, err := http.Post("http://localhost:8080/RouteP", "application/json",  bytes.NewBuffer([]byte("{\"Message\": \"Test\"}")))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	ByteBuffer, err := ioutil.ReadAll(res.Body)
	got := string(ByteBuffer)
	log.Println("response data: ",got)

	want := `{"Message": "Post Method Only!", "Auth": "true"}`

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}

}


func TestRoutePG(t *testing.T) {

	res, err := http.Get("http://localhost:8080/RouteP")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	ByteBuffer, err := ioutil.ReadAll(res.Body)
	got := string(ByteBuffer)
	log.Println("response data: ",got)

	want := `{"Message": "Method is Not Valid", "Auth": "false"}`

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}

}


// func BenchmarkRoute1(b *testing.B) {
// 	for i :=0; i < b.N ; i++ {
// 		// RunFunction Here
// 	}
// }