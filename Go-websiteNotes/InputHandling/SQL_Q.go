package main

import (
	"fmt"
	"log"
)

func SQLInjParse(query string) bool{
	// testing for bugs

	var Ichar = []string{"=","-",";",":","'"}

	var i = 0
	for i < len(Ichar){
		res1 := strings.Contains(query, Ichar[i])
		if res1 == true{
			return false
		}else{
			return true
		}
		i++
	}
	// res1 := strings.Contains(query, "=")
 //    res2 := strings.Contains(query, "-")
 //    res3 := strings.Contains(query, ";")
 //    res4 := strings.Contains(query, ":")
 //    res5 := strings.Contains(query, "'")

 //    if res1 == true{
 //    	return false
 //    }else if res2 == true{
 //    	return false
 //    }else if res3 == true{
 //    	return false
 //    }else if res4 == true{
 //    	return false
 //    } else if res5 == true{
 //    	return false
 //    }else {
 //    	return true
 //    }

}

func Home(w http.ResponseWriter, r *http.Request) {
	log.Print("Running Home Page\n")

}

func main() {

	http.HandleFunc("/", Home)
	log.Print("Listening.....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}