package main


import (
	"database/sql"
	"fmt"
	// "strings"
	"net/http"
	"html/template"
	"log"
	// "time"

	_ "github.com/go-sql-driver/mysql"
)


var tpl *template.Template

type OGs struct {
    NFT_id string // could be int 
    SS string
    Skin string
    Visors string
    Eyes string
    OnEyes string
    mouth string
    Hats string
    chains string
    BK string
}

type Apes struct {
    NFT_id string // could be int 
    SS string
    Skin string
    Visors string
    Eyes string
    OnEyes string
    mouth string
    Tail string
    chains string
    BK string
}

type Pups struct {
    NFT_id string // could be int 
    SS string
    Skin string
    Visors string
    Eyes string
    OnEyes string
    mouth string
    Earrings string
    chains string
    BK string
}


func AstroOGs(w http.ResponseWriter, r *http.Request) {


	data := OGs{
	    NFT_id: r.FormValue("NFT_Id"),
	    SS: r.FormValue("SS"),
	    Skin: r.FormValue("Skin"),
	    Visors: r.FormValue("Visors"),
	    Eyes: r.FormValue("Eyes"),
	    OnEyes: r.FormValue("OnEyes"),
	    mouth: r.FormValue("mouth"),
	    Hats: r.FormValue("Hats"),
	    chains: r.FormValue("chains"),
	    BK: r.FormValue("BK"),
	}
	_ = data

	fmt.Println(data)
	if data.NFT_id != ""{
		Q := SearchID(data.NFT_id, 0)
		fmt.Println(Q)
	}



	tpl.ExecuteTemplate(w, "OGs.html", "null")
}

func AstroApes(w http.ResponseWriter, r *http.Request) {

	data := Apes{
	    NFT_id: r.FormValue("NFT_Id"),
	    SS: r.FormValue("SS"),
	    Skin: r.FormValue("Skin"),
	    Visors: r.FormValue("Visors"),
	    Eyes: r.FormValue("Eyes"),
	    OnEyes: r.FormValue("OnEyes"),
	    mouth: r.FormValue("mouth"),
	    Tail: r.FormValue("Tail"),
	    chains: r.FormValue("chains"),
	    BK: r.FormValue("BK"),
	}
	_ = data

	fmt.Println(data)

	tpl.ExecuteTemplate(w, "Apes.html", "null")
}

func AstroPups(w http.ResponseWriter, r *http.Request) {

	data := Pups{
	    NFT_id: r.FormValue("NFT_Id"),
	    SS: r.FormValue("SS"),
	    Skin: r.FormValue("Skin"),
	    Visors: r.FormValue("Visors"),
	    Eyes: r.FormValue("Eyes"),
	    OnEyes: r.FormValue("OnEyes"),
	    mouth: r.FormValue("mouth"),
	    Earrings: r.FormValue("Earrings"),
	    chains: r.FormValue("chains"),
	    BK: r.FormValue("BK"),
	}
	_ = data

	fmt.Println(data)

	tpl.ExecuteTemplate(w, "Pups.html", "null")
	
}

func Index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", "null")
}

func SearchID(NFT_Id string, collection int) []string {
	var Coll string

	if collection == 0 {
		Coll = "OGs"
	}else if collection == 1 {
		Coll = "Apes"
	}else if collection == 2 {
		Coll = "Pups"
	}

	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE astrosearch")
	if err != nil {
	log.Fatal(err)
	}

	var (

	    ID string 
	    Suit string
	    skin string
	    Visor string
	    Eye string
	    oneyes string
	   	Mouth string
	    CTrait string
	    Chains string
	    bk string
	)

	query, err := db.Query(`SELECT NFT_ID FROM ? WHERE NFT_ID = ?`, Coll, NFT_Id)
	if err != nil {
		log.Fatal(err)
	}
	defer query.Close()

	for query.Next() {
		err := query.Scan(&ID, &Suit, &skin, &Visor, &Eye, &oneyes, &Mouth, &CTrait, &Chains, &bk)
		if err != nil {
			log.Fatal(err)
		}
	}

	NFT := []string{ID, Suit, skin, Visor, Eye, oneyes, Mouth, CTrait, Chains, bk}

	return NFT

}

// func SearchTraits(data , collection int) {
// 	var Coll string

// 	if collection == 0 {
// 		Coll = "OGs"
// 	}else if collection == 1 {
// 		Coll = "Apes"
// 	}else if collection == 2 {
// 		Coll = "Pups"
// 	}
	
// 	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := db.Ping(); err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = db.Exec("USE astrosearch")
// 	if err != nil {
// 	log.Fatal(err)
// 	}

// 	var (

// 	    ID string 
// 	    Suit string
// 	    skin string
// 	    Visor string
// 	    Eye string
// 	    oneyes string
// 	   	Mouth string
// 	    CTrait string
// 	    chains string
// 	    bk string
// 	)

// 	// query the database for all NFTs with traits from data
	

// 	fmt.Println(data)
// }



func main() {

	tpl, _ = template.ParseGlob("./static/Templates/*html")

	http.HandleFunc("/", Index)
	http.HandleFunc("/AstroOGs", AstroOGs)
	http.HandleFunc("/AstroApes", AstroApes)
	http.HandleFunc("/AstroPups", AstroPups)

	log.Print("Listening....")
	err := http.ListenAndServeTLS(":9000", "localhost.crt", "localhost.key", nil)
	if err != nil {
			log.Fatal("ListenAndServe: ", err)
	}

}
