package xrp

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestXrpWallet_FromGenerate(t *testing.T) {
	wallet := XrpWallet{}
	wallet.Test = false
	wallet.FromGenerate()
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	//4a15295773437550dc5333804c1252bc60d8fda405fc1a1a43888b0f1a0e970c
	//033e4d9adadb4d4619e0b47d58f8aef9fc40a135726fb833d3c51001c25ea7af64
	//place educate street cycle zero gasp fat gap actual tortoise concert pottery
	//r3SFaE2xqGwGEfq92JVw8MeJZ9k4ML2viF
}

func TestXrpWallet_FromPrivateKey(t *testing.T) {
	wallet := XrpWallet{}
	wallet.Test = false
	wallet.FromPrivateKey("4a15295773437550dc5333804c1252bc60d8fda405fc1a1a43888b0f1a0e970c")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "4a15295773437550dc5333804c1252bc60d8fda405fc1a1a43888b0f1a0e970c", wallet.GetPrivateKey())
	assert.Equal(t, "033e4d9adadb4d4619e0b47d58f8aef9fc40a135726fb833d3c51001c25ea7af64", wallet.GetPublicKey())
	assert.Equal(t, "r3SFaE2xqGwGEfq92JVw8MeJZ9k4ML2viF", wallet.GetAddress())
}

func TestXrpWallet_FromMnemonic(t *testing.T) {
	wallet := XrpWallet{}
	wallet.Test = false
	mnemonic := "siege resist defy forum seven find creek audit invite marine favorite civil"
	wallet.FromMnemonic(mnemonic)
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "3b8c63d5792b8150b884e5a56d94e39edca96dc61b6e95b53d9c2bdff58f0030", wallet.GetPrivateKey())
	assert.Equal(t, "031ca8c01c4c37ad9c3fcacec54918024657e405b80bc679acae096cf05487b598", wallet.GetPublicKey())
	assert.Equal(t, "rDtbkVwYCmz7137nvJJNja3dkNuhjJmK7u", wallet.GetAddress())
}

func TestXrpWallet_FromMnemonicAndPath(t *testing.T) {
	wallet := XrpWallet{}
	wallet.Test = false
	mnemonic := "siege resist defy forum seven find creek audit invite marine favorite civil"
	wallet.FromMnemonicAndPath(mnemonic, "m/44'/144'/0'/0/0")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	assert.Equal(t, "3b8c63d5792b8150b884e5a56d94e39edca96dc61b6e95b53d9c2bdff58f0030", wallet.GetPrivateKey())
	assert.Equal(t, "031ca8c01c4c37ad9c3fcacec54918024657e405b80bc679acae096cf05487b598", wallet.GetPublicKey())
	assert.Equal(t, "rDtbkVwYCmz7137nvJJNja3dkNuhjJmK7u", wallet.GetAddress())
}

func TestXrpWallet_FromPublicKey(t *testing.T) {
	wallet := XrpWallet{}
	wallet.Test = false
	wallet.FromPublicKey("031ca8c01c4c37ad9c3fcacec54918024657e405b80bc679acae096cf05487b598")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "rDtbkVwYCmz7137nvJJNja3dkNuhjJmK7u", wallet.GetAddress())
}
