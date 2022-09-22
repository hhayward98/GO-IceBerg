package WebStruct

import (
	"log"
	"net/http"
	"html/template"
)

// WebStruct is read into each template as the data.
// modify and add objects where needed within WebStruct

// Header, Body, footer are all arrays 

type WebStruct struct {
	Head Header
	Body body
	Foot Footer
}


// Domain name is used for routing links correctly on web app.
type Header struct {
	DomainName string	// http://example.com/8080 
	Statictemp string	// /static/templates/index.html
	// Modify for what is needed.	
}


// an Array of somthing 
type body struct {
	Blist []struct
}

// credits in the footer 
type Footer struct {
	Credits string
}


func CreateTemplateStruct() {
	H := Header{
		DomainName: "http://localhost:8080/",
		StaticFiles: "/static/templates/",
	}

	B := Body{
		Blist: []
	}

	F := Footer{
		Credits: "Create By Aesir Constructs"
	}

	webstruct := WebStruct{
		Head: H
		Body: B
		Foot: F
	}
	_ = webstruct

	return webstruct
}





