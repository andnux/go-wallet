package go_wallet

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-crypto"
	"strconv"
	"strings"
)

type FileCoinWallet struct {
}

func (wallet *FileCoinWallet) Name() string {
	return "FIL"
}

func (wallet *FileCoinWallet) Broadcast(data []byte) string {
	return ""
}

func (wallet *FileCoinWallet) Signature(data []byte, privateKey string) []byte {
	return []byte{}
}

func (wallet *FileCoinWallet) Generate(test bool) string {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return wallet.GenerateByMnemonic(mnemonic, "m/44'/461'/0'/0/0", test)
}
func (wallet *FileCoinWallet) GenerateByPrivateKey(privateKey string, test bool) string {
	keyHex, err := hex.DecodeString(privateKey)
	if err != nil {
		fmt.Println(err)
	}
	publicKey := crypto.PublicKey(keyHex)
	k1Address, err := address.NewSecp256k1Address(publicKey)
	if err != nil {
		fmt.Println(err)
	}
	a := Account{
		PrivateKey: privateKey,
		PublicKey:  hex.EncodeToString(publicKey),
		Address:    k1Address.String()}
	res, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	return string(res)
}

func (wallet *FileCoinWallet) GenerateByMnemonic(mnemonic string, path string, test bool) string {
	s := strings.Split(path, "/")
	parseUint, err := strconv.ParseUint(strings.ReplaceAll(s[len(s)-1], "'", ""), 0, 1)
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
	keyFromMnemonic, err := bip44.NewKeyFromMnemonic(mnemonic, 0x800001cd, uint32(account), uint32(chain), uint32(parseUint))
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	key := keyFromMnemonic.Key
	publicKey := crypto.PublicKey(key)
	k1Address, err := address.NewSecp256k1Address(publicKey)
	if err != nil {
		fmt.Println(err)
	}
	a := Account{
		PrivateKey: hex.EncodeToString(key),
		PublicKey:  hex.EncodeToString(publicKey),
		Mnemonic:   mnemonic,
		Address:    k1Address.String()}
	res, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	return string(res)
}
