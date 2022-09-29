package main

import (
	"log"
	"fmt"
	"crypto/rsa"
	"crypto/rand"
)



func GenKeys() {

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}


	publicKey := privateKey.PublicKey


}

func Encrypt(msg string) {

// sign and validate Hash
	Bmsg := []byte(msg)
	msgHash := sha256.New() 
	_, err = msgHash.Write(msg)
	if err != nil {
		log.Fatal(err)
	}
	msgHshSum, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = rsa.VerifyPSS(&publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		fmt.Println("could not verify signature: ", err)
		return
	}



// encrypt
	encryptedBytes, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, &publicKey, []byte("super secret message"), nil))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Encrypted message : ", encryptedBytes)
}

func Decrypt() {

	decryptedBytes, err := privateKey.Decrypt(nil, encryptedBytes, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Decrypted message: ", string(decryptedBytes))

}



