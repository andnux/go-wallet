package go_wallet

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTronGenerate(t *testing.T) {
	wallet := InitTronWallet(false)
	account := wallet.Generate()
	//{"private_key":"f5fcfe50f3d5f148b6749a1ef32237bf529e658421bded1c57e7207d0f938d62",
	//"public_key":"02c837f5683142d22c2b430b3ed2a2ae66ed454828e8480dbf14f6e1320fd2352d",
	//"address":"THqz6SSHRRjoKRzx7CF3GsVRBWjH957gbht",
	//"mnemonic":"salute wrestle level autumn village fiber spend organ arm urge citizen save",
	//"keystore":""}
	fmt.Println(account)
}

func TestTronGenerateByPrivateKey(t *testing.T) {
	wallet := InitTronWallet(false)
	privateKey := "c11259d1227eaa6b555338d433184a06eee5001c0855360b25d4c3ae4cb2c2f8"
	account := wallet.GenerateByPrivateKey(privateKey)
	fmt.Println(account)
	//publicKey := "048af5ab1f19ec29f72464ff49f9df0f375202f27d7f68c655a02f317f90d529431105ccb96bfe0f86a72014519059966fa0483996358be18b669041255ad70ab6"
	address := "TCfE3nLLrgtwL4tusA6AX4WcsszKCDkW5P"
	var a Account
	err := json.Unmarshal([]byte(account), &a)
	if err != nil {
		fmt.Println(err)
	}
	//assert.Equal(t, publicKey, a.PublicKey)
	assert.Equal(t, address, a.Address)
}

func TestTronGenerateByMnemonic(t *testing.T) {
	wallet := InitTronWallet(false)
	mnemonic := "soccer abuse buyer upset calm pass extra camp visa man economy elephant"
	account := wallet.GenerateByMnemonic(mnemonic, "m/44'/195'/0'/0/0")
	fmt.Println(account)
	publicKey := "Tron7KrBYdccTUYxExpJiXDp2YXMjMxJHepKade93gxhuHw5PTmx9k"
	var a Account
	err := json.Unmarshal([]byte(account), &a)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, publicKey, a.PublicKey)
}
