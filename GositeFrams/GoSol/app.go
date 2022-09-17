package main

import (
	"fmt"
	"log"
	"net/http"
	"context"
	"html/template"

	"github.com/portto/solana-go-sdk/client"
	// "github.com/portto/solana-go-sdk/client/rpc"
	"github.com/portto/solana-go-sdk/types"


)

var tpl *template.Template

var C = client.NewClient(rpc.MainnetRPCEndpoint)

type Wallet struct {
	account types.Account
	C 	*client.Client
}


func ErrorHandler(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ImportWallet(privateKey []byte, RPCEndpoint string) (Wallet, error) {

	wallet, err := types.AccountFromBytes(privateKey)
	if err != nil {
		return Wallet{}, err
	}

	return Wallet{
		wallet,
		client.NewClient(RPCEndpoint),
	}, nil
}

func WalletSearch(w http.ResponseWriter, r *http.Request) {


	walletID := r.FormValue("WalletID")

	balance, err := C.GetBalance(context.TODO(),walletID)

	ErrorHandler(err)

	fmt.Println("Wallet balance in Lamport:", balance)
	fmt.Println("Wallet balance in SOL:", balance/1e9)
	tpl.ExecuteTemplate(w, "index.html", "")

}

func main() {


	tpl, _ = template.ParseGlob("./static/templates/*html")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", WalletSearch)

	log.Print("Listening.....")
	log.Fatal(http.ListenAndServe(":8080", nil))

}


