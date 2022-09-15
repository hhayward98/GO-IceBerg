package ASearch

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"

)

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



func SearchTraits(data []string , collection int) {

	var NFTList []
	
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
	    chains string
	    bk string
	)

	if collection == 0 {

		query, err := db.Query(`SELECT * FROM OGs WHERE Suit, skin, Visor, Eye, oneyes, Mouth, CTrait, chains, bk = (?,?,?,?,?,?,?,?,?)`, data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8])
		if err != nil {
			log.Fatal(err)
		}
		defer query.Close()

		for query.Next() {
			err := query.Scan(&ID, &Suit, &skin, &Visor, &Eye, &oneyes, &Mouth, &CTrait, &chains, &bk)
			if err != nil {
				log.Fatal(err)
			}
		}


	}else if collection == 1 {

		query, err := db.Query(`SELECT * FROM OGs WHERE Suit, skin, Visor, Eye, oneyes, Mouth, CTrait, chains, bk = (?,?,?,?,?,?,?,?,?)`, data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8])
		if err != nil {
			log.Fatal(err)
		}
		defer query.Close()

		for query.Next() {
			err := query.Scan(&ID, &Suit, &skin, &Visor, &Eye, &oneyes, &Mouth, &CTrait, &chains, &bk)
			if err != nil {
				log.Fatal(err)
			}
		}

	}else if collection == 2 {

		query, err := db.Query(`SELECT * FROM OGs WHERE Suit, skin, Visor, Eye, oneyes, Mouth, CTrait, chains, bk = (?,?,?,?,?,?,?,?,?)`, data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7], data[8])
		if err != nil {
			log.Fatal(err)
		}
		defer query.Close()

		for query.Next() {
			err := query.Scan(&ID, &Suit, &skin, &Visor, &Eye, &oneyes, &Mouth, &CTrait, &chains, &bk)
			if err != nil {
				log.Fatal(err)
			}
		}

	}
	// query the database for all NFTs with traits from data
	

	fmt.Println(data)
}