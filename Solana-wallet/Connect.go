package main 


import (
	"context"
	"fmt"
	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/portto/solana-go-sdk/types"
)

var C := client.NewClient(rpc.MainnetRPCEndpoint)

type Wallet struct {
	account types.Account
	C 	*client.Client
}

func CreateWallet(RPCEndpoint string) Wallet {

	return Wallet{
		types.NewAccount(),
		client.NewClient(RPCEndpoint),
	}
}


func GetBalence() {

	Bal, err := C.GetBalance(context.TODO(), "Wallet-address")

	if err != nil {
		panic(err)
	}


}

func IWallet(privateKey []byte, RPCEndpoint string) (Wallet, error) {
	wallet, err := types.AccountFromBytes(privateKey)
	if err != nil {
		return Wallet{}, err
	}

	return Wallet{
		wallet,
		client.NewClient(RPCEndpoint),
	}, nil
}

func ImportWallet() {
	// import wallet with base58
	Base58, err := types.AccountFromBase58("")
	if err != nil {
		panic(err)
	}
	// import wallet with bytes slice private key

	ByteSlice, err := types.AccountFromBytes([]byte{})
	if err != nil {
		panic(err)
	}
	// import wallet with hex private key
	HexKey, err := types.AccountFromHex("")
	if err != nil {
		panic(err)
	}
}



func main() {

	// C := client.NewClient(rpc.MainnetRPCEndpoint)

	res, err := C.GetVersion(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Println("version", res.SolanaCore)
}