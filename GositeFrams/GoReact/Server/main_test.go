package main 

import (
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

func TestRoute2(t *testing.T) {

	res, err := http.Get("http://localhost:8080/RouteTwo")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	ByteBuffer, err := ioutil.ReadAll(res.Body)
	got := string(ByteBuffer)
	log.Println("response data: ",got)

	want := ""

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

	want := ""

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

	want := ""

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}

}

func TestRouteP(t *testing.T) {

	res, err := http.Get("http://localhost:8080/RouteP")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	ByteBuffer, err := ioutil.ReadAll(res.Body)
	got := string(ByteBuffer)
	log.Println("response data: ",got)

	want := ""

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}

}


// func BenchmarkRoute1(b *testing.B) {
// 	for i :=0; i < b.N ; i++ {
// 		// RunFunction Here
// 	}
// }