package go_wallet

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileCoinGenerate(t *testing.T) {
	wallet := InitFileCoinWallet(false)
	account := wallet.Generate(false)
	//{"private_key":"a7b339a280908c4308b1c25e70ce3d78d8c95ac68ee8a032b489ec84fb5add44",
	//"public_key":"04cc1f598b5671da6aa6416bfcfdb1923ac30ccf7b7cab732adedc7095a4840ab163dfa04a492e3443c0b17044517da9bd180f492e61d12db9db742ae554c0cc25",
	//"address":"t1bgz6unoi2zl3e2bv3nsngalnrqt36cqnhvvgxeq",
	//"mnemonic":"crash spatial blade carry blush stand donkey trade field razor sweet grow",
	//"keystore":""}
	fmt.Println(account)
}

func TestFileCoinGenerateByPrivateKey(t *testing.T) {
	wallet := InitFileCoinWallet(false)
	privateKey := "a7b339a280908c4308b1c25e70ce3d78d8c95ac68ee8a032b489ec84fb5add44"
	account := wallet.GenerateByPrivateKey(privateKey)
	fmt.Println(account)
	var a Account
	err := json.Unmarshal([]byte(account), &a)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, "04cc1f598b5671da6aa6416bfcfdb1923ac30ccf7b7cab732adedc7095a4840ab163dfa04a492e3443c0b17044517da9bd180f492e61d12db9db742ae554c0cc25", a.PublicKey)
	assert.Equal(t, "t1bgz6unoi2zl3e2bv3nsngalnrqt36cqnhvvgxeq", a.Address)
}

func TestFileCoinGenerateByMnemonic(t *testing.T) {
	wallet := InitFileCoinWallet(false)
	mnemonic := "crash spatial blade carry blush stand donkey trade field razor sweet grow"
	account := wallet.GenerateByMnemonic(mnemonic, "m/44'/461'/0'/0/0")
	fmt.Println(account)
	var a Account
	err := json.Unmarshal([]byte(account), &a)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, "04cc1f598b5671da6aa6416bfcfdb1923ac30ccf7b7cab732adedc7095a4840ab163dfa04a492e3443c0b17044517da9bd180f492e61d12db9db742ae554c0cc25", a.PublicKey)
	assert.Equal(t, "t1bgz6unoi2zl3e2bv3nsngalnrqt36cqnhvvgxeq", a.Address)
}
