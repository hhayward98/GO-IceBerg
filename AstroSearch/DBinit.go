package main 

import (
	"database/sql"
	// "encoding/json"
	"log"
	"fmt"
	// "io/ioutil"
	"os"
	"bufio"

	"strings"


	_ "github.com/go-sql-driver/mysql"
)




func main() {

	//connect to database
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

	// {
	// 	query := `
	// 		CREATE TABLE OGIMG (
	// 			id INT NOT NULL,
	// 			imgSrc TEXT NOT NULL,
	// 			PRIMARY KEY (id));`

	// 	if _, err := db.Exec(query); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// {
	// 	query := `
	// 		CREATE TABLE OGs (
	// 		    id INT NOT NULL,
	// 		    Suit TEXT NOT NULL,
	// 		    skin TEXT NOT NULL,
	// 		    Visor TEXT NOT NULL,
	// 		    Eye TEXT NOT NULL,
	// 		    oneyes TEXT NOT NULL,
	// 		   	Mouth TEXT NOT NULL,
	// 		    CTrait TEXT NOT NULL,
	// 		    chains TEXT NOT NULL,
	// 		    bk TEXT NOT NULL,
	// 		    PRIMARY KEY (id));`

	// 	if _, err := db.Exec(query); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// {
	// 	query2 := `
	// 		CREATE TABLE Apes (
	// 		    id INT NOT NULL,
	// 		    Suit TEXT NOT NULL,
	// 		    skin TEXT NOT NULL,
	// 		    Visor TEXT NOT NULL,
	// 		    Eye TEXT NOT NULL,
	// 		    oneyes TEXT NOT NULL,
	// 		   	Mouth TEXT NOT NULL,
	// 		    CTrait TEXT NOT NULL,
	// 		    chains TEXT NOT NULL,
	// 		    bk TEXT NOT NULL,
	// 		    PRIMARY KEY (id));`
		
	// 	if _, err := db.Exec(query2); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	// {
	// 	query3 := `
	// 		CREATE TABLE Pups (
	// 			id INT NOT NULL,
	// 		    Suit TEXT NOT NULL,
	// 		    skin TEXT NOT NULL,
	// 		    Visor TEXT NOT NULL,
	// 		    Eye TEXT NOT NULL,
	// 		    oneyes TEXT NOT NULL,
	// 		   	Mouth TEXT NOT NULL,
	// 		    CTrait TEXT NOT NULL,
	// 		    chains TEXT NOT NULL,
	// 		    bk TEXT NOT NULL,
	// 		    PRIMARY KEY (id));`
		
	// 	if _, err := db.Exec(query3); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }


	// =======================================================================================

	readFile, err := os.Open("OGsIMG.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var Sray []string
	for fileScanner.Scan() {
		Sray = append(Sray, fileScanner.Text())

	}


	for i, s := range Sray {
		sl := strings.Split(s, ",")
		fmt.Println(i, sl)
	
		//remove s from https
		// HTTP := strings.Replace(sl[1], "s", "", 1)


		//====================================================================================

		// result, err := db.Exec(`INSERT IGNORE INTO ogimg (id, imgSrc) Values(?,?)`, sl[0], HTTP)
		// if err != nil {
		// 	log.Fatal(err)
		// }
		// id, err := result.LastInsertId()
		// fmt.Println(id)

		// =================================================================================

		// 	// 0		1  		2 		3  		4 				5		6 				7 		8  			 9 		10       	11 		12 			13 			14 			15 		16 			17 			18 
		// 	// [3793, Visors, No Visor, Hats, Two Side Hair, On Eyes, Yellow Goggles, Mouth, Mad Max Paint, Eyes, Blue Eyes, Chains, No Chain, Space Suits, Tattooed Suit, Skins, Zebra Skin, Backgrounds, Orange BG ]
	

		// if len(sl) > 12{

		// 	result, err := db.Exec(`INSERT IGNORE INTO OGs (id, Suit, skin, Visor, Eye, oneyes, Mouth, CTrait, chains, bk) Values(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, sl[0], sl[14], sl[16], sl[2], sl[10], sl[6], sl[8], sl[4], sl[12], sl[18])
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	id, err := result.LastInsertId()
		// 	fmt.Println(id)

		// }

		// // Sl is a list of traits for each NFT with its id number at the start
		// // load sl in the database under the correct table

		// // ==============================================================

		var SRC string
		var ID int 
		DBcheck, _ := db.Query(`SELECT id,imgSrc FROM OGIMG WHERE id = ?`, sl[0])
		defer DBcheck.Close()
		for DBcheck.Next() {
			err := DBcheck.Scan(&ID, &SRC)
			if err != nil{
				log.Fatal(err)
			}
		}

		fmt.Println(SRC)

	}

	readFile.Close()


	// =========================================================================================================



	// load 
	// {
	// 	query := `DROP TABLE ogimg;`
	// 	if _, err := db.Exec(query); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	// {

	// Create Table for image URL with NFT_id as the primary key (maybe)
	// 



}




// ========================
// 		Structure
// ========================


// Database: Asearch
// 
// 	Tables: 

// 		OGs: {
// 		    id INT NOT NULL,
// 		    Suit TEXT NOT NULL,
// 		    skin TEXT NOT NULL,
// 		    Visor TEXT NOT NULL,
// 		    Eye TEXT NOT NULL,
// 		    oneyes TEXT NOT NULL,
// 		   	Mouth TEXT NOT NULL,
// 		    CTrait TEXT NOT NULL,
// 		    chains TEXT NOT NULL,
// 		    bk TEXT NOT NULL,
// 		    PRIMARY KEY (id));
//		}

//		Apes: {
// 		    id INT NOT NULL,
// 		    Suit TEXT NOT NULL,
// 		    skin TEXT NOT NULL,
// 		    Visor TEXT NOT NULL,
// 		    Eye TEXT NOT NULL,
// 		    oneyes TEXT NOT NULL,
// 		   	Mouth TEXT NOT NULL,
// 		    CTrait TEXT NOT NULL,
// 		    chains TEXT NOT NULL,
// 		    bk TEXT NOT NULL,
// 		    PRIMARY KEY (id));
//		}


//		Pups: {
// 		    id INT NOT NULL,
// 		    Suit TEXT NOT NULL,
// 		    skin TEXT NOT NULL,
// 		    Visor TEXT NOT NULL,
// 		    Eye TEXT NOT NULL,
// 		    oneyes TEXT NOT NULL,
// 		   	Mouth TEXT NOT NULL,
// 		    CTrait TEXT NOT NULL,
// 		    chains TEXT NOT NULL,
// 		    bk TEXT NOT NULL,
// 		    PRIMARY KEY (id));
//		}

// 		OG-img: {
// 			id INT NOT NULL,
// 			imgSrc TEXT NOT NULL,
// 			PRIMARY KEY (id));
// 		}
// 		Apes-img: {
// 			id INT NOT NULL,
// 			imgSrc TEXT NOT NULL,
// 			PRIMARY KEY (id));
// 		}
// 		Pups-img: {
// 			id INT NOT NULL,
// 			imgSrc TEXT NOT NULL,
// 			PRIMARY KEY (id));
// 		}

