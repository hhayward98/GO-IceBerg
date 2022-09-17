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

type SBI_NFT struct {
    NFT_id string 
    SS string
    Skin string
    Visors string
    Eyes string
    OnEyes string
    mouth string
    Hats string
    chains string
    BK string
    imgsrc string
}

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

	var NFT []string
	var SearchByID bool

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
		SearchByID = true
		Q := SearchID(data.NFT_id, 0)
		fmt.Println(Q)
		NFT = Q
		var FoundNFT = SBI_NFT{
		    NFT_id: NFT[0], 
		    SS: NFT[1],
		    Skin: NFT[2],
		    Visors: NFT[3],
		    Eyes: NFT[4],
		    OnEyes: NFT[5],
		    mouth: NFT[6],
		    Hats: NFT[7],
		    chains: NFT[8],
		    BK: NFT[9],
		    imgsrc: "",
		}

		tpl.ExecuteTemplate(w, "OGs.html", FoundNFT)
		return
	} else {
		SearchByID = false
		Tlist := []string{data.SS, data.Skin, data.Visors, data.Eyes, data.OnEyes, data.mouth, data.Hats, data.chains, data.BK}
		// Q := SearchTraits(Tlist, 0)
		fmt.Println(Tlist)
	}
	// fmt.Println(NFT[0])

	fmt.Println(SearchByID)

// if SearchID = true then use SBI_NFT struct and as the data input for ExecuteTemplate
// need to query database table OG-imgSrc for image src of NFT and add to SBI_NFT struct along with other NFT-metadata

	tpl.ExecuteTemplate(w, "OGs.html", NFT)
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
	if data.NFT_id != "" {
		Q := SearchID(data.NFT_id, 1)
		
		fmt.Println(Q)
	} else {
		Tlist := []string{data.SS, data.Skin, data.Visors, data.Eyes, data.OnEyes, data.mouth, data.Tail, data.chains, data.BK}
		// Q := SearchTraits(Tlist, 1)
		fmt.Println(Tlist)
	}

	tpl.ExecuteTemplate(w, "Apes.html", "null")
	return
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
	if data.NFT_id != "" {
		Q := SearchID(data.NFT_id, 2)
		fmt.Println(Q)
	} else {
		Tlist := []string{data.SS, data.Skin, data.Visors, data.Eyes, data.OnEyes, data.mouth, data.Earrings, data.chains, data.BK}
		// Q := SearchTraits(Tlist, 2)
		fmt.Println(Tlist)
	}

	tpl.ExecuteTemplate(w, "Pups.html", "null")
	return
	
}

func Index(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("static/templates/index.html"))

	tmpl.Execute(w, nil)
	return
}

func SearchID(NFT_Id string, collection int) []string {

	var NFT []string

	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE asearch")
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

	if collection == 0 {
		query, err := db.Query(`SELECT * FROM OGs WHERE id = ?`, NFT_Id)
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

	} else if collection == 1 {
		query, err := db.Query(`SELECT * FROM Apes WHERE id = ?`, NFT_Id)
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

	} else if collection == 2 {

		query, err := db.Query(`SELECT * FROM Pups WHERE id = ?`, NFT_Id)
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

	return NFT

}

// func SearchTraits(data []string , collection int) {


	
// 	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := db.Ping(); err != nil {
// 		log.Fatal(err)
// 	}

// 	_, err = db.Exec("USE asearch")
// 	if err != nil {
// 		log.Fatal(err)
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

// 	if collection == 0 {

// 		query, err := db.Query(`SELECT * FROM OGs WHERE Suit, skin, Visor, Eye, oneyes, Mouth, CTrait, chains, bk = (?,?,?,?,?,?,?,?,?)`, data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer query.Close()

// 		for query.Next() {
// 			err := query.Scan(&ID, &Suit, &skin, &Visor, &Eye, &oneyes, &Mouth, &CTrait, &chains, &bk)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		}


// 	}else if collection == 1 {

// 		query, err := db.Query(`SELECT * FROM OGs WHERE Suit, skin, Visor, Eye, oneyes, Mouth, CTrait, chains, bk = (?,?,?,?,?,?,?,?,?)`, data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer query.Close()

// 		for query.Next() {
// 			err := query.Scan(&ID, &Suit, &skin, &Visor, &Eye, &oneyes, &Mouth, &CTrait, &chains, &bk)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		}

// 	}else if collection == 2 {

// 		query, err := db.Query(`SELECT * FROM OGs WHERE Suit, skin, Visor, Eye, oneyes, Mouth, CTrait, chains, bk = (?,?,?,?,?,?,?,?,?)`, data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		defer query.Close()

// 		for query.Next() {
// 			err := query.Scan(&ID, &Suit, &skin, &Visor, &Eye, &oneyes, &Mouth, &CTrait, &chains, &bk)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		}

// 	}
// 	// query the database for all NFTs with traits from data
	

// 	fmt.Println(data)
// }



func main() {

	tpl, _ = template.ParseGlob("./static/templates/*html")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	
	http.HandleFunc("/", Index)
	http.HandleFunc("/AstroOGs", AstroOGs)
	http.HandleFunc("/AstroApes", AstroApes)
	http.HandleFunc("/AstroPups", AstroPups)

	log.Print("Listening....")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
