package main

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "strings"

    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/common"

    token "zeptest/zep_token"
)


func main() {
    key := `{"address":"36a892d4990b8ed4d9669b6b49a358f9ec7846e2","crypto":{"cipher":"aes-128-ctr","ciphertext":"10180f16480f4adec08b2df395d7a1647108369984c2286637a7b1b27e623833","cipherparams":{"iv":"2339ed05d32ab6ef2540c3869850a92c"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"bd648b30f41dfb8aee1fc0f0d86505bea95535852b6e26003e60c972ac21ed10"},"mac":"b58f3aa4d78a11a8b9cee792b85ae4da1f5ad36802ac833f9db6e0481451e712"},"id":"e8aa5d18-0483-4725-bea5-8a769b44b5eb","version":3}`

    client, err := ethclient.Dial("http://0.0.0.0:22000")
    if err != nil {
        log.Fatal(err)
    }

    fromAddres := common.HexToAddress("0x36a892d4990b8ed4d9669b6b49a358f9ec7846e2")
    nonce, err := client.PendingNonceAt(context.Background(), fromAddres)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Nonce: ", nonce)

    auth, err := bind.NewTransactor(strings.NewReader(key), "secret")
    if err != nil {
        log.Fatalf("Failed to create authorized transactor: %v", err)
    }

    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("SuggestGasPrice: ", gasPrice)

    auth.Nonce = big.NewInt(int64(nonce + 2))
    auth.Value = big.NewInt(2300000)     // in wei
    auth.GasLimit =  uint64(3000000)     // in units
    auth.GasPrice = big.NewInt(0)

    address, tx, instance, err := token.DeployToken(auth, client)
    if err != nil {
        log.Fatal("Deployment Error: ", err)
    }

    fmt.Println(address.Hex())
    fmt.Println(tx.Hash().Hex())

    name, err := instance.Name(nil)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(name)
}