package go_wallet

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVsysGenerate(t *testing.T) {
	wallet := VsysWallet{}
	account := wallet.Generate(false)
	//{"private_key":"5KQNYR9UDPeALTNGqYxmyzDNJeiqEyHbUrpwT5tC5Wgg",
	//"public_key":"F5saYo9gDHuqXBRbkpWB1pgdj3HWR3SXaLCD7hYMhUem",
	//"address":"ARHBWHd15nx7i9oGk91bMYgTas6Y6qtCTW2",
	//"mnemonic":"clog lady equip lens ensure ladder lava diamond report kick rabbit cook",
	//"keystore":""}
	fmt.Println(account)
}

func TestVsysGenerateByPrivateKey(t *testing.T) {
	wallet := VsysWallet{}
	privateKey := "5KQNYR9UDPeALTNGqYxmyzDNJeiqEyHbUrpwT5tC5Wgg"
	account := wallet.GenerateByPrivateKey(privateKey, false)
	fmt.Println(account)
	var a Account
	err := json.Unmarshal([]byte(account), &a)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, "F5saYo9gDHuqXBRbkpWB1pgdj3HWR3SXaLCD7hYMhUem", a.PublicKey)
	assert.Equal(t, "ARHBWHd15nx7i9oGk91bMYgTas6Y6qtCTW2", a.Address)
}

func TestVsysGenerateByMnemonic(t *testing.T) {
	wallet := VsysWallet{}
	mnemonic := "clog lady equip lens ensure ladder lava diamond report kick rabbit cook"
	account := wallet.GenerateByMnemonic(mnemonic, "m/44'/360'/0'/0/0", false)
	fmt.Println(account)
	var a Account
	err := json.Unmarshal([]byte(account), &a)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, "F5saYo9gDHuqXBRbkpWB1pgdj3HWR3SXaLCD7hYMhUem", a.PublicKey)
	assert.Equal(t, "ARHBWHd15nx7i9oGk91bMYgTas6Y6qtCTW2", a.Address)
}
