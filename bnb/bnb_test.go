package bnb

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestBnbWallet_BuildFromRandomGenerate(t *testing.T) {
	wallet := BnbWallet{}
	wallet.Test = false
	wallet.BuildFromRandomGenerate()
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	//b26cbd62ee55220a372fe80fd7abc35611871e5812dd7becbb13a88372240d70
	//eb5ae9872102e07e605340b1a22bfdd460459bf6104dbc005a9960703e720744922d7be5e573
	//ladder jungle old medal spoil limit sibling crystal apology snap sauce issue
	//bnb15qfl6722ajyl9qhr8v3srt9m3rvlvud8mach8u
}

func TestBnbWallet_BuildFromPrivateKey(t *testing.T) {
	wallet := BnbWallet{}
	wallet.Test = false
	wallet.BuildFromPrivateKey("b26cbd62ee55220a372fe80fd7abc35611871e5812dd7becbb13a88372240d70")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "bnb15qfl6722ajyl9qhr8v3srt9m3rvlvud8mach8u", wallet.GetAddress())
}

func TestBnbWallet_BuildFromMnemonic(t *testing.T) {
	wallet := BnbWallet{}
	wallet.Test = false
	mnemonic := "notable gap make parrot tail deposit desert oxygen blanket crunch sense village"
	wallet.BuildFromMnemonic(mnemonic)
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "bnb1xuqxrwsvru64xfe7r220pjprqdhel6sujf6s93", wallet.GetAddress())
}

func TestBnbWallet_BuildFromMnemonicAndPath(t *testing.T) {
	wallet := BnbWallet{}
	wallet.Test = false
	mnemonic := "undo dynamic dust become chat cage pool junk sphere next rent creek"
	wallet.BuildFromMnemonicAndPath(mnemonic, "m/44'/714'/0'/0/0")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	//452dc601aaa5a5910ceccc145908b922c66794c58a1afec80048966471b1aadb
	//eb5ae987210284332b2e15a66f0154383c2b9a8e4282ca6b2f55c8730cdf9b81a749a861a3ff
	//undo dynamic dust become chat cage pool junk sphere next rent creek
	//bnb13yltnnh68qp4pxa2x0mju7rvyy2vp7la889qch
}

func TestBnbWallet_BuildFromPublicKey(t *testing.T) {
	wallet := BnbWallet{}
	wallet.Test = false
	wallet.BuildFromPublicKey("eb5ae987210284332b2e15a66f0154383c2b9a8e4282ca6b2f55c8730cdf9b81a749a861a3ff")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "bnb13yltnnh68qp4pxa2x0mju7rvyy2vp7la889qch", wallet.GetAddress())
}
