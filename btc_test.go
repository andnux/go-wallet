package go_wallet

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBtcGenerate(t *testing.T) {
	wallet := BtcWallet{}
	account := wallet.Generate(false)
	//{"private_key":"0x4845353de4fc654e9abc45c2c79b360470f882231a1e117f600da2ce11660930",
	//"public_key":"0x04f0089738bcb6073dd5ec30cee040ae5460e42de95a4c14e7bf01fde72fe86e1e13726b317a5ddc202c9be8157659b9e4903885c49aada90f526d5dfe635e2ca3",
	//"address":"178x1Un8tJebjw1F7YrrzNiVa6quHqKxVd",
	//"mnemonic":"when anger shock gain armed bird cup virus kite hybrid mix citizen",
	//"keystore":""}
	fmt.Println(account)
}

func TestBtcGenerateByPrivateKey(t *testing.T) {
	wallet := BtcWallet{}
	privateKey := "0x4845353de4fc654e9abc45c2c79b360470f882231a1e117f600da2ce11660930"
	account := wallet.GenerateByPrivateKey(privateKey, false)
	fmt.Println(account)
	publicKey := "0x04f0089738bcb6073dd5ec30cee040ae5460e42de95a4c14e7bf01fde72fe86e1e13726b317a5ddc202c9be8157659b9e4903885c49aada90f526d5dfe635e2ca3"
	address := "178x1Un8tJebjw1F7YrrzNiVa6quHqKxVd"
	var a Account
	err := json.Unmarshal([]byte(account), &a)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, publicKey, a.PublicKey)
	assert.Equal(t, address, a.Address)
}

func TestBtcGenerateByMnemonic(t *testing.T) {
	wallet := BtcWallet{}
	mnemonic := "when anger shock gain armed bird cup virus kite hybrid mix citizen"
	account := wallet.GenerateByMnemonic(mnemonic, "m/44'/0'/0'/0/0", false)
	fmt.Println(account)
	publicKey := "0x04f0089738bcb6073dd5ec30cee040ae5460e42de95a4c14e7bf01fde72fe86e1e13726b317a5ddc202c9be8157659b9e4903885c49aada90f526d5dfe635e2ca3"
	address := "178x1Un8tJebjw1F7YrrzNiVa6quHqKxVd"
	var a Account
	err := json.Unmarshal([]byte(account), &a)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, publicKey, a.PublicKey)
	assert.Equal(t, address, a.Address)
}
