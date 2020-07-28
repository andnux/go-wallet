package go_wallet

import (
	"encoding/hex"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
)

type TemplateWallet struct {
	Test       bool
	privateKey *string
	publicKey  *string
	mnemonic   *string
	address    *string
}

func (wallet *TemplateWallet) Sign(data []byte) (signed []byte) {
	if wallet.privateKey == nil {
		return signed
	}
	bytes, err := hex.DecodeString(*wallet.privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes)
	return []byte{}
}

func (wallet *TemplateWallet) BuildFromRandomGenerate() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.BuildFromMnemonic(mnemonic)
}

func (wallet *TemplateWallet) BuildFromPrivateKey(hexKey string) {
	wallet.publicKey = nil
	wallet.privateKey = &hexKey
	wallet.mnemonic = nil
	bytes, err := hex.DecodeString(*wallet.privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes)
}

func (wallet *TemplateWallet) GetPrivateKey() string {
	key := wallet.privateKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *TemplateWallet) publicKeyToAddress(hexPublicKey string) string {
	bytes, err := hex.DecodeString(hexPublicKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes)
	return ""
}

func (wallet *TemplateWallet) BuildFromPublicKey(publicKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.publicKey = &publicKey
	addr := wallet.publicKeyToAddress(publicKey)
	wallet.address = &addr
}

func (wallet *TemplateWallet) GetPublicKey() string {
	key := wallet.publicKey
	if key == nil {
		return ""
	}
	return *key
}
func (wallet *TemplateWallet) BuildFromMnemonic(mnemonic string) {
	wallet.BuildFromMnemonicAndPath(mnemonic, "m/44'/145'/0'/0/0")
}

func (wallet *TemplateWallet) BuildFromMnemonicAndPath(mnemonic string, path string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	parser, err := bip44Parser(path)
	if err != nil {
		panic(err)
	}
	key, err := bip44.NewKeyFromMnemonic(mnemonic,
		parser.CoinType, parser.Account,
		parser.Change, parser.AddressIndex)
	if err != nil {
		panic(err)
	}
	fmt.Println(key.Key)
}

func (wallet *TemplateWallet) GetMnemonic() string {
	if wallet.mnemonic == nil {
		return ""
	}
	return *wallet.mnemonic
}

func (wallet *TemplateWallet) BuildFromAddress(address string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.address = &address
}

func (wallet *TemplateWallet) GetAddress() string {
	if wallet.address == nil {
		return ""
	}
	return *wallet.address
}
