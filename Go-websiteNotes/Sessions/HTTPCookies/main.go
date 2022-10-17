package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
)


/* Cookie Struct 

type Cookie struct {
	Name string 
	Value string

	path string 	- indicates a URL path that must exist in request URL, to send Cookie Header.
	Domain string 	- specifies which hosts are allowed to recive the cookie
	Expires time.Time - deletes at specified date
	RawExpires string - for reading cookies only

	// maxAge=0 : no Max-Age attribute specified
	// maxAge<0 : delete cookie now
	// maxAge>0 : 

	MaxAge int 		- deletes cookie after specified amount of time.
	Secure bool 	- sent to a server on encrypted request, never unsecured(HTTP)
	HttpOnly bool 	- not accessible by JavaScript, only sent to server.
	SameSite SameSite - servers requier that a cookie shouldn't be sent with cross-origin request
	Raw string
	Unparsed []string - raw text of unparsed attribute-value paris
}

*/

type Cookie struct {
	Name string
	Value string

	// Path string
	// Domain string
	// Expires time.Time
	// RawExpires string

	// MaxAge int 
	// Secure bool
	HttpOnly bool
	// SameSite SameSite
	// Raw string
	// Unparsed []string
}


var tpl *template.Template

func Home(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("1st-cookie")
	fmt.Println("cookie:", cookie, "err:", err)
	if err != nil {
		fmt.Println("cookie was not found")
		cookie = &http.Cookie{
			Name: "1st-cookie",
			Value: "my First cookie value",
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
	}

	// test
	IpAddy := r.Header.Get("X-Real-Ip")
	if IpAddy == "" {
		IpAddy = r.Header.Get("X-Forwarded-For")
	}
	if IpAddy == "" {
		IpAddy = r.RemoteAddr
	}
	// prints in IP format Ex: 127.0.0.1 
	fmt.Println("IPAddress : ",IpAddy)


	tpl.ExecuteTemplate(w, "index.html", nil)

}




func main() {
	tpl, _ = template.ParseGlob("templates/*.html")
	

	http.HandleFunc("/", Home)
	log.Println("Listening.....")
	log.Fatal(http.ListenAndServe(":8080", nil))


}