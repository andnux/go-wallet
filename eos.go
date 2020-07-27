package go_wallet

import (
	"errors"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/eoscanada/eos-go/ecc"
)

type EosWallet struct {
	Test       bool
	privateKey *ecc.PrivateKey
	publicKey  *ecc.PublicKey
	mnemonic   *string
}

func (wallet *EosWallet) Sign(data []byte) (signed []byte, err error) {
	if wallet.privateKey == nil {
		return signed, errors.New("请先导入私钥")
	}
	out, err := wallet.privateKey.Sign(data)
	if err != nil {
		fmt.Println(err)
	}
	return out.Content, nil
}

func (wallet *EosWallet) BuildFromRandomGenerate() {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.BuildFromMnemonic(mnemonic)
}

func (wallet *EosWallet) BuildFromPrivateKey(privateKey string) {
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

func (wallet *EosWallet) BuildFromPublicKey(publicKey string) {
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

func (wallet *EosWallet) BuildFromMnemonic(mnemonic string) {
	wallet.BuildFromMnemonicAndPath(mnemonic, "m/44'/194'/0'/0/0")
}

func (wallet *EosWallet) BuildFromMnemonicAndPath(mnemonic string, path string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	parser, err := bip44Parser(path)
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
