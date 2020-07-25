package go_wallet

import (
	"encoding/json"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/eoscanada/eos-go/ecc"
	"strconv"
	"strings"
)

type EosWallet struct {
}

func (wallet *EosWallet) Name() string {
	return "EOS"
}

func (wallet *EosWallet) Broadcast(data []byte) string {
	return ""
}

func (wallet *EosWallet) Signature(data []byte, privateKey string) []byte {
	return []byte{}
}

func (wallet *EosWallet) Generate(test bool) string {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	key, err := bip44.NewKeyFromMnemonic(mnemonic, 0x80000019, 0, 0, 0)
	if err != nil {
		fmt.Println(err)
	}
	privateKey, err := ecc.NewPrivateKeyFromSeed(string(key.Key))
	if err != nil {
		fmt.Println(err)
	}
	publicKey := privateKey.PublicKey().String()
	account := Account{PrivateKey: privateKey.String(),
		PublicKey: publicKey,
		Mnemonic:  mnemonic}
	bytes, err := json.Marshal(account)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}
func (wallet *EosWallet) GenerateByPrivateKey(privateKey string, test bool) string {
	key, err := ecc.NewPrivateKey(privateKey)
	if err != nil {
		fmt.Println(err)
	}
	publicKey := key.PublicKey().String()
	account := Account{PrivateKey: key.String(), PublicKey: publicKey}
	bytes, err := json.Marshal(account)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}

func (wallet *EosWallet) GenerateByMnemonic(mnemonic string, path string, test bool) string {
	s := strings.Split(path, "/")
	address, err := strconv.ParseUint(strings.ReplaceAll(s[len(s)-1], "'", ""), 0, 1)
	if err != nil {
		fmt.Println(err)
	}
	chain, err := strconv.ParseUint(strings.ReplaceAll(s[len(s)-2], "'", ""), 0, 1)
	if err != nil {
		fmt.Println(err)
	}
	account, err := strconv.ParseUint(strings.ReplaceAll(s[len(s)-3], "'", ""), 0, 1)
	if err != nil {
		fmt.Println(err)
	}
	key, err := bip44.NewKeyFromMnemonic(mnemonic, 0x80000019, uint32(account), uint32(chain), uint32(address))
	if err != nil {
		fmt.Println(err)
	}
	privateKey, err := ecc.NewPrivateKeyFromSeed(string(key.Key))
	if err != nil {
		fmt.Println(err)
	}
	publicKey := privateKey.PublicKey().String()
	a := Account{PrivateKey: privateKey.String(),
		PublicKey: publicKey,
		Mnemonic:  mnemonic}
	bytes, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}
