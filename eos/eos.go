package eos

import (
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/andnux/go-wallet"
	"github.com/eoscanada/eos-go/ecc"
)

type EosWallet struct {
	Test       bool
	BaseUrl    string
	privateKey *ecc.PrivateKey
	publicKey  *ecc.PublicKey
	mnemonic   *string
}

func (wallet *EosWallet) Sign(data []byte) (signed []byte) {
	if wallet.privateKey == nil {
		return signed
	}
	out, err := wallet.privateKey.Sign(data)
	if err != nil {
		fmt.Println(err)
	}
	return out.Content
}

func (wallet *EosWallet) FromGenerate() {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.FromMnemonic(mnemonic)
}

func (wallet *EosWallet) FromPrivateKey(privateKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	key, err := ecc.NewPrivateKey(privateKey)
	if err != nil {
		fmt.Println(err)
	}
	publicKey := key.PublicKey()
	wallet.publicKey = &publicKey
	wallet.privateKey = key
}

func (wallet *EosWallet) GetPrivateKey() string {
	key := wallet.privateKey
	if key == nil {
		return ""
	}
	return key.String()
}

func (wallet *EosWallet) FromPublicKey(publicKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	key, err := ecc.NewPublicKey(publicKey)
	if err != nil {
		fmt.Println(err)
	}
	wallet.publicKey = &key
}

func (wallet *EosWallet) GetPublicKey() string {
	key := wallet.publicKey
	if key == nil {
		return ""
	}
	return key.String()
}

func (wallet *EosWallet) FromMnemonic(mnemonic string) {
	wallet.FromMnemonicAndPath(mnemonic, "m/44'/194'/0'/0/0")
}

func (wallet *EosWallet) FromMnemonicAndPath(mnemonic string, path string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	parser, err := go_wallet.Bip44Parser(path)
	if err != nil {
		fmt.Println(err)
	}
	key, err := bip44.NewKeyFromMnemonic(mnemonic,
		parser.CoinType, parser.Account,
		parser.Change, parser.AddressIndex)
	if err != nil {
		fmt.Println(err)
	}
	privateKey, err := ecc.NewPrivateKeyFromSeed(string(key.Key))
	if err != nil {
		fmt.Println(err)
	}
	publicKey := privateKey.PublicKey()
	wallet.publicKey = &publicKey
	wallet.privateKey = privateKey
	wallet.mnemonic = &mnemonic
}

func (wallet *EosWallet) GetMnemonic() string {
	if wallet.mnemonic == nil {
		return ""
	}
	return *wallet.mnemonic
}
