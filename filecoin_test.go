package go_wallet

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileCoinGenerate(t *testing.T) {
	wallet := InitFileCoinWallet(false)
	account := wallet.Generate()
	//{"private_key":"a7b339a280908c4308b1c25e70ce3d78d8c95ac68ee8a032b489ec84fb5add44",
	//"public_key":"04cc1f598b5671da6aa6416bfcfdb1923ac30ccf7b7cab732adedc7095a4840ab163dfa04a492e3443c0b17044517da9bd180f492e61d12db9db742ae554c0cc25",
	//"address":"t1bgz6unoi2zl3e2bv3nsngalnrqt36cqnhvvgxeq",
	//"mnemonic":"crash spatial blade carry blush stand donkey trade field razor sweet grow",
	//"keystore":""}
	//{"private_key":"ddfe3848f9753c32fafb40fd8da3ee7bbbc32915db3f1ade75933ed8a24a4f64",
	//"public_key":"0422a6f48a58698155db86086caa3e042be304e4ef25fc858acf14710ef30c63a760a631ea209f50c1a0ec1a9edad0d675404271615ef33a0c8ae7a3384a02ec4c",
	//"address":"f1v77px6n2vwsum5f76u6wqh4wpiwuw6ow4byviui",
	//"mnemonic":"food forward empty chest best agree east matter hill rebel merry hub",
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
