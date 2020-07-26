package go_wallet

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/andnux/go-wallet/address"
	"github.com/andnux/go-wallet/crypto"
	"strconv"
	"strings"
)

type FileCoinWallet struct {
	Test bool
}

func (wallet *FileCoinWallet) Name() string {
	return "FIL"
}

func (wallet *FileCoinWallet) Signature(data []byte, privateKey string) []byte {
	keyHex, err := hex.DecodeString(privateKey)
	if err != nil {
		fmt.Println(err)
	}
	sign, err := crypto.Sign(keyHex, data)
	if err != nil {
		fmt.Println(err)
	}
	return sign
}

func (wallet *FileCoinWallet) Generate() string {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return wallet.GenerateByMnemonic(mnemonic, "m/44'/461'/0'/0/0")
}
func (wallet *FileCoinWallet) GenerateByPrivateKey(privateKey string) string {
	keyHex, err := hex.DecodeString(privateKey)
	if err != nil {
		fmt.Println(err)
	}
	publicKey := crypto.PublicKey(keyHex)
	k1Address, err := address.NewSecp256k1Address(publicKey)
	if err != nil {
		fmt.Println(err)
	}
	var network address.Network
	if wallet.Test {
		network = address.Testnet
	} else {
		network = address.Mainnet
	}
	a := Account{
		PrivateKey: privateKey,
		PublicKey:  hex.EncodeToString(publicKey),
		Address:    k1Address.String(network)}
	res, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	return string(res)
}

func (wallet *FileCoinWallet) GenerateByMnemonic(mnemonic string, path string) string {
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
	var network address.Network
	if wallet.Test {
		network = address.Testnet
	} else {
		network = address.Mainnet
	}
	a := Account{
		PrivateKey: hex.EncodeToString(key),
		PublicKey:  hex.EncodeToString(publicKey),
		Mnemonic:   mnemonic,
		Address:    k1Address.String(network)}
	res, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	return string(res)
}
