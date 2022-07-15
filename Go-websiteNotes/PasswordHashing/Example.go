/*Password Hashing*/

package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "super-secret-key"
	hash, _ := HashPassword(password) // should check for error

	fmt.Println("Password:", password)
	fmt.Println("Hash: 	", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match: 	", match)

}