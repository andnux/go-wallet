package go_wallet

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEthGenerate(t *testing.T) {
	wallet := EthWallet{}
	account := wallet.Generate(false)
	//{"private_key":"a6b8d4da2f530cbb023c17dffd8629a1eb6efa438c355ded93350e3114e560cf",
	//"public_key":"027b6fa83246908856e4c5a2956b4860d3baef2b7950deb86487e5d0f758f4b9f4",
	//"address":"0x794efafb6e19b9465befe4082db599fa73ed9cad",
	//"mnemonic":"admit blossom boring smoke chicken category narrow fuel deliver butter weekend vanish",
	//"keystore":""}
	fmt.Println(account)
}

func TestEthGenerateByPrivateKey(t *testing.T) {
	wallet := EthWallet{}
	privateKey := "a6b8d4da2f530cbb023c17dffd8629a1eb6efa438c355ded93350e3114e560cf"
	account := wallet.GenerateByPrivateKey(privateKey, false)
	fmt.Println(account)
	var a Account
	err := json.Unmarshal([]byte(account), &a)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, "027b6fa83246908856e4c5a2956b4860d3baef2b7950deb86487e5d0f758f4b9f4", a.PublicKey)
	assert.Equal(t, "0x794efafb6e19b9465befe4082db599fa73ed9cad", a.Address)
}

func TestEthGenerateByMnemonic(t *testing.T) {
	wallet := EthWallet{}
	mnemonic := "admit blossom boring smoke chicken category narrow fuel deliver butter weekend vanish"
	account := wallet.GenerateByMnemonic(mnemonic, "m/44'/60'/0'/0/0", false)
	fmt.Println(account)
	var a Account
	err := json.Unmarshal([]byte(account), &a)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, "027b6fa83246908856e4c5a2956b4860d3baef2b7950deb86487e5d0f758f4b9f4", a.PublicKey)
	assert.Equal(t, "0x794efafb6e19b9465befe4082db599fa73ed9cad", a.Address)
}
