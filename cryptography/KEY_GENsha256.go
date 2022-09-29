package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	var Uput string

	fmt.Println("Enter a string for encoding")
	fmt.Scanln(&Uput)

	Ecrpt := sha256.Sum256([]byte(Uput))
	fmt.Println("encoding string....\n")
	fmt.Printf("%x",Ecrpt)
}