package main

import (
	"strings"

	"log"
	"net/http"
	"html/template"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)


var tpl *template.Template

type HTMLDATA struct {
	DomainName string
	Message string
	Body BODY
	footerContent string
}

type BODY struct{
	HeroList []*SuperHuman
	VillainList []*SuperHuman
}


type SuperHuman struct {
	Name string
	PassiveP string
	AttackP string
}


func SetHTMLData() HTMLDATA{


	bodyData := BODY{
		HeroList: make([]*SuperHuman,0),
		VillainList: make([]*SuperHuman,0),
	}

	Hdata := HTMLDATA{
		DomainName: "http://localhost:8080",
		Message: " ",
		Body: bodyData,
		footerContent: "Created by <a href='https://demonic-labs.com'>Demonic Labs</a>",
	}

	return Hdata
}


func Debugger(err error, Etype int) {
	if err != nil {
		if Etype == 0 {
			log.Fatal(err)
		}else if Etype == 1 {
			log.Println("=================================================")
			log.Println(err)
			log.Println("==================================================")
		}
	}
}


func InjectionHandler(Uput string) bool {

	res1 := strings.Contains(Uput, "=")
    res2 := strings.Contains(Uput, "-")
    res3 := strings.Contains(Uput, ";")
    res4 := strings.Contains(Uput, "<")
    res5 := strings.Contains(Uput, "'")
    res6 := strings.Contains(Uput, "`")

    if res1 == true{
    	return false
    }else if res2 == true{
    	return false
    }else if res3 == true{
    	return false
    }else if res4 == true{
    	return false
    } else if res5 == true{
    	return false
    }else if res6 == true {
    	return false
    }else {
    	return true
    }


}

func RequestFromDatabase(Allegiance string) []*SuperHuman{

	// db, err := sql.Open("mysql", "test:toor@tcp(db:3306)/superhumans")
	// Debugger(err, 1)
	// if err := db.Ping(); err != nil {
	// 	log.Fatal(err)
	// }

	// return list of super humans based on Allegince.

	// Test DBConnect
	db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("USE superhumans")
	if err != nil {
		log.Fatal(err)
	}


	if Allegiance == "Hero" {

		superheros := make([]*SuperHuman,0)

		heroRows, err := db.Query(`SELECT heroname, passivepower, attackpower FROM heros`)
		Debugger(err, 0)

		for heroRows.Next() {
			Hero := new(SuperHuman)
			if err := heroRows.Scan(&Hero.Name, &Hero.PassiveP, &Hero.AttackP); err != nil {
				Debugger(err, 0)
			}

			superheros = append(superheros, Hero)
		}

		if err := heroRows.Err(); err != nil {
			Debugger(err, 0)
		}
		heroRows.Close()

		//superheros is an array of SuperHumans from the heros table
		return superheros


	} else if Allegiance == "Villain" {

		supervillains := make([]*SuperHuman,0)

		VillainRows, _ := db.Query(`SELECT villainname, passivepower, attackpower FROM villains`)
		Debugger(err, 0)

		for VillainRows.Next() {
			Villain := new(SuperHuman)
			
			if err := VillainRows.Scan(&Villain.Name, &Villain.PassiveP, &Villain.AttackP); err != nil {
				Debugger(err, 0)
			}
			
			supervillains = append(supervillains, Villain)

		}

		if err := VillainRows.Err(); err != nil {
			Debugger(err, 0)
		}

		VillainRows.Close()

		return supervillains

	}
	return nil
}






func Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Running Home Page....")

	htlmData := SetHTMLData()

	tpl.ExecuteTemplate(w, "index.html", htlmData)
	return

}


func AddSuperHuman(w http.ResponseWriter, r *http.Request) {
	log.Println("Running AddSuperHuman Page....")
	


	htlmData := SetHTMLData()



	if r.Method != http.MethodPost {
		tpl.ExecuteTemplate(w, "addsuperhuman.html", htlmData)
		return
	} else if r.Method == "POST" {

		// docker SQL db connect
		// db, err := sql.Open("mysql", "Test:toor@tcp(db:3306)/superhumans")
		// Debugger(err, 1)
		// if err := db.Ping(); err != nil {
		// 	log.Fatal(err)
		// }


		// Testing 
		db, err := sql.Open("mysql", "Test:toor@(127.0.0.1:3308)/?parseTime=true")
		if err != nil {
			log.Fatal(err)
		}
		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}

		_, err = db.Exec("USE superhumans")
		if err != nil {
			log.Fatal(err)
		}

		log.Println("Connected to Database")


		Allegiance := r.FormValue("allegiance")

		SuperhumanName := r.FormValue("SHN")

		PassivePower := r.FormValue("Pasv")

		AttackPower := r.FormValue("Attk")

		// sanitize user input

		if len(SuperhumanName) < 1 {
			log.Println("Empty Field")
			htlmData.Message = "Name can not be Empty!"
			tpl.ExecuteTemplate(w, "addsuperhuman.html", htlmData)
			return
		}

		if len(PassivePower) < 1 {
			log.Println("Empty Field")
			htlmData.Message = "Passive Power can not be Empty!"
			tpl.ExecuteTemplate(w, "addsuperhuman.html", htlmData)
			return
		}

		if len(AttackPower) < 1 {
			log.Println("Empty Field")
			htlmData.Message = "Attack Power can not be Empty!"
			tpl.ExecuteTemplate(w, "addsuperhuman.html", htlmData)
			return
		}

		if InjectionHandler(SuperhumanName) != true {
			log.Print("invalid characters detected!!")
			htlmData.Message = "Illegal characters detected!!"
			tpl.ExecuteTemplate(w, "addsuperhuman.html", htlmData)
			return
		}

		if InjectionHandler(PassivePower) != true {
			log.Print("invalid characters detected!!")
			htlmData.Message = "Illegal characters detected!!"
			tpl.ExecuteTemplate(w, "addsuperhuman.html", htlmData)
			return
		}

		if InjectionHandler(AttackPower) != true {
			log.Print("invalid characters detected!!")
			htlmData.Message = "Illegal characters detected!!"
			tpl.ExecuteTemplate(w, "addsuperhuman.html", htlmData)
			return
		}

		Namelower := strings.ToLower(SuperhumanName)


		if Allegiance == "Hero" {


			var SelcetedHero string
			HeroCheck, _ := db.Query(`SELECT heroname FROM heros WHERE heroname = ?`, Namelower)

			for HeroCheck.Next() {
				err := HeroCheck.Scan(&SelcetedHero)
				Debugger(err, 0)
			}
			HeroCheck.Close()



			if len(SelcetedHero) > 0 {
				log.Println("SuperHero Name not available")

				htlmData.Message = "SuperHero Name not available"

				tpl.ExecuteTemplate(w, "addsuperhuman.html", htlmData)
				return


			} else {
				log.Println("Adding SuperHero to Database")
				result, err := db.Exec(`INSERT INTO heros (heroname, passivepower, attackpower) VALUES (?,?,?)`, Namelower, PassivePower, AttackPower)
				Debugger(err,0)
				log.Println("Insert successful: ",result)
				htlmData.Message = "Successfully added superHuman!"

				tpl.ExecuteTemplate(w, "addsuperhuman.html", htlmData)
				return
			}



		} else if Allegiance == "Villain" {
			var SelectedVillain string
			VillainCheck, _ := db.Query(`SELECT villainname FROM villains WHERE villainname = ?`, Namelower)
			
			for VillainCheck.Next() {
				err := VillainCheck.Scan(&SelectedVillain)
				Debugger(err, 0)
			}

			VillainCheck.Close()

			if len(SelectedVillain) > 0 {

				htlmData.Message = "Super Villain name is not available"

				tpl.ExecuteTemplate(w, "addsuperhuman.html", htlmData)
				return

			} else {
				log.Println("Adding SuperVillain to Database")
				result, err := db.Exec(`INSERT INTO villains (villainname, passivepower, attackpower) VALUES (?,?,?)`, Namelower, PassivePower, AttackPower)
				Debugger(err, 0)
				log.Println("Successfully added SuperVillain to Database", result)
				htlmData.Message = "Successfully Uploaded SuperVillain!"

				tpl.ExecuteTemplate(w, "addsuperhuman.html", htlmData)
				return

			}

		} else {

			log.Println("ELSE: ******************************************************")
			htlmData.Message = "Make Sure all fields and options and not empty"
			tpl.ExecuteTemplate(w, "addsuperhuman.html", htlmData)

		}


	}


}


func ShowHero(w http.ResponseWriter, r *http.Request) {
	log.Println("Running ShowHero Page....")


	ArrayHeros := RequestFromDatabase("Hero")

	log.Println(ArrayHeros)

	htlmData := SetHTMLData()

	htlmData.Body.HeroList = ArrayHeros

	tpl.ExecuteTemplate(w, "showhero.html", htlmData)
	return


}


func ShowVillains(w http.ResponseWriter, r *http.Request) {
	log.Println("Running ShowVillains page....")


	ArryVillains := RequestFromDatabase("Villain")
	
	log.Println(ArryVillains)

	htlmData := SetHTMLData()

	htlmData.Body.VillainList = ArryVillains

	tpl.ExecuteTemplate(w, "showvillain.html", htlmData)
	return


}

func Battle(w http.ResponseWriter, r *http.Request) {
	log.Println("Running Battle Page....")


	ArryVillains := RequestFromDatabase("Villain")
	ArrayHeros := RequestFromDatabase("Hero")

	htlmData := SetHTMLData()
	htlmData.Body.VillainList = ArryVillains
	htlmData.Body.HeroList = ArrayHeros

	tpl.ExecuteTemplate(w, "battle.html", htlmData)

	return

}

func AppRoutes() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", Home)
	http.HandleFunc("/AddSuper", AddSuperHuman)
	http.HandleFunc("/ShowHeros", ShowHero)
	http.HandleFunc("/ShowVillains", ShowVillains)
	http.HandleFunc("/SuperBattle", Battle)


	log.Fatal(http.ListenAndServe(":8080", nil))


}


func main() {

	tpl, _ = template.ParseGlob("./static/templates/*.html")

	log.Println("Listening....")

	AppRoutes()


}