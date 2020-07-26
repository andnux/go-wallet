package go_wallet

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBtcGenerate(t *testing.T) {
	wallet := InitBtcWallet(false)
	account := wallet.Generate()
	//{"private_key":"5JzHoCgTjJjhrNUwZbAfPtQCWPZrJJCRcxX2h9gTUBSFaB6D4ZQ",
	//"public_key":"04a8e62453e8f6d3f29d39acc9ba64588d6ae836f3bedbbb8677696167d3654747e1abf4f7628b1a7f75d403f59c17d6d8df2c965320209fb8e82b5bbb227edfa4",
	//"address":"1KkfsD58VqXKxmsuyvPe4dNCJC9DPjACHk",
	//"mnemonic":"nose ski weekend heavy ozone spring limit salon ask bread lift window",
	//"keystore":""}
	fmt.Println(account)
}

func TestBtcGenerateByPrivateKey(t *testing.T) {
	wallet := InitBtcWallet(false)
	privateKey := "5JzHoCgTjJjhrNUwZbAfPtQCWPZrJJCRcxX2h9gTUBSFaB6D4ZQ"
	account := wallet.GenerateByPrivateKey(privateKey)
	fmt.Println(account)
	publicKey := "04a8e62453e8f6d3f29d39acc9ba64588d6ae836f3bedbbb8677696167d3654747e1abf4f7628b1a7f75d403f59c17d6d8df2c965320209fb8e82b5bbb227edfa4"
	address := "1KkfsD58VqXKxmsuyvPe4dNCJC9DPjACHk"
	var a Account
	err := json.Unmarshal([]byte(account), &a)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, publicKey, a.PublicKey)
	assert.Equal(t, address, a.Address)
}

func TestBtcGenerateByMnemonic(t *testing.T) {
	wallet := InitBtcWallet(false)
	mnemonic := "nose ski weekend heavy ozone spring limit salon ask bread lift window"
	account := wallet.GenerateByMnemonic(mnemonic, "m/44'/0'/0'/0/0")
	fmt.Println(account)
	publicKey := "04a8e62453e8f6d3f29d39acc9ba64588d6ae836f3bedbbb8677696167d3654747e1abf4f7628b1a7f75d403f59c17d6d8df2c965320209fb8e82b5bbb227edfa4"
	address := "1KkfsD58VqXKxmsuyvPe4dNCJC9DPjACHk"
	var a Account
	err := json.Unmarshal([]byte(account), &a)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, publicKey, a.PublicKey)
	assert.Equal(t, address, a.Address)
}
