package main

import (
	"fmt"
	"strings"

)

type Test struct {
	HTML string
	Body string
	foot string
}

func SLInjParse(query string) bool{
	// testing for bugs

	var Ichar = []string{"=","-",";",":","'","`"}


	fmt.Println(query)
	for i, obj := range Ichar {
		fmt.Println(i, obj)
		res1 := strings.Contains(query, obj)
		if res1 == true{
			// when res1 returns true the function loop ends 
			return false
		}else{
			// so when res1 dose not contain the first value in list then function returns before checking others
			return true
		}
	}
	return false

}


func main() {

	var P = Test{HTML: "temp", Body: "", foot: "null"}
	fmt.Println(P)
	P.Body = "Testing"
	fmt.Println("starting...")
	fmt.Println(P)
	T1 := SLInjParse("hello")
	T2 := SLInjParse("WHILE 1 = 1, DELETE * FROM users")
	T3 := SLInjParse("Lol'1+1=2';")
	T4 := SLInjParse("1`create database`;")
	T5 := SLInjParse(":32")

	booLL := []bool{T1, T2, T3, T4, T5}

	for i, obj := range booLL {
		fmt.Println(i, obj)
	}


}