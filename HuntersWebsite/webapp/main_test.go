package main 

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	// "github.com/rs/cors"

)



func TestGetHome(t *testing.T) {

	res, err := http.Get("http://localhost:8080/Home")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	// ByteBuffer, err := ioutil.ReadAll(res.Body)

	// got := string(ByteBuffer)

	// log.Println("response data: ",got)

	// want := 

	// if got != want {
	// 	t.Errorf("got %q, wanted %q", got , want)
	// }

}


func TestPostHome(t *testing.T) {

	res, err := http.Post("http://localhost:8080/Home", "application/json",  bytes.NewBuffer([]byte("{\"Message\": \"Test\"}")))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	ByteBuffer, err := ioutil.ReadAll(res.Body)

	got := string(ByteBuffer)

	// log.Println("response data: ",got)

	want := "Method Not Allowed\n"

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}

}

func TestPostProjects(t *testing.T) {


	res, err := http.Post("http://localhost:8080/Projects", "application/json",  bytes.NewBuffer([]byte("{\"Message\": \"Test\"}")))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	ByteBuffer, err := ioutil.ReadAll(res.Body)

	got := string(ByteBuffer)

	// log.Println("response data: ",got)

	want := "Method Not Allowed\n"

	if got != want {
		t.Errorf("got %q, wanted %q", got , want)
	}


}


func TestPostSkills(t *testing.T) {

	res, err := http.Post("http://localhost:8080/Skills", "application/json", bytes.NewBuffer([]byte("{\"Message\": \"Test\"}")))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	ByteBuffer, err := ioutil.ReadAll(res.Body)

	got := string(ByteBuffer)

	want := "Method Not Allowed\n"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	
	}

	
}