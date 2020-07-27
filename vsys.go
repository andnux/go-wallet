package go_wallet

import (
	"errors"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/walkbean/vsys-sdk-go/vsys"
)

type VsysWallet struct {
	Test       bool
	privateKey *string
	publicKey  *string
	mnemonic   *string
	address    *string
}

func (wallet *VsysWallet) Sign(data []byte) (signed []byte, err error) {
	if wallet.privateKey == nil {
		return signed, errors.New("请先导入私钥")
	}
	var account *vsys.Account
	if wallet.Test {
		account = vsys.InitAccount(vsys.Testnet)
	} else {
		account = vsys.InitAccount(vsys.Mainnet)
	}
	account.BuildFromPrivateKey(*wallet.privateKey)
	return []byte(account.SignData(data)), nil
}

func (wallet *VsysWallet) BuildFromRandomGenerate() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.BuildFromMnemonic(mnemonic)
}

func (wallet *VsysWallet) BuildFromPrivateKey(privateKey string) {
	var account *vsys.Account
	if wallet.Test {
		account = vsys.InitAccount(vsys.Testnet)
	} else {
		account = vsys.InitAccount(vsys.Mainnet)
	}
	account.BuildFromPrivateKey(privateKey)
	wallet.privateKey = &privateKey
	publicKey := account.PublicKey()
	wallet.publicKey = &publicKey
	address := account.Address()
	wallet.address = &address
}

func (wallet *VsysWallet) GetPrivateKey() string {
	key := wallet.privateKey
	if key == nil {
		return ""
	}
	return *key
}
func (wallet *VsysWallet) BuildFromPublicKey(publicKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.publicKey = &publicKey
	if wallet.Test {
		address := vsys.PublicKeyToAddress(publicKey, vsys.Testnet)
		wallet.address = &address
	} else {
		address := vsys.PublicKeyToAddress(publicKey, vsys.Mainnet)
		wallet.address = &address
	}
}

func (wallet *VsysWallet) GetPublicKey() string {
	key := wallet.publicKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *VsysWallet) BuildFromMnemonic(mnemonic string) {
	wallet.BuildFromMnemonicAndPathAndNonce(mnemonic, "m/44'/360'/0'/0/0", 0)
}

func (wallet *VsysWallet) BuildFromMnemonicAndPath(mnemonic string, path string) {
	wallet.BuildFromMnemonicAndPathAndNonce(mnemonic, path, 0)
}

func (wallet *VsysWallet) BuildFromMnemonicAndPathAndNonce(mnemonic string, path string, nonce int) {
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
	var vsysAccount *vsys.Account
	if wallet.Test {
		vsysAccount = vsys.InitAccount(vsys.Testnet)
	} else {
		vsysAccount = vsys.InitAccount(vsys.Mainnet)
	}
	vsysAccount.BuildFromSeed(string(key.Key), nonce)
	privateKey := vsysAccount.PrivateKey()
	wallet.privateKey = &privateKey
	publicKey := vsysAccount.PublicKey()
	wallet.publicKey = &publicKey
	address := vsysAccount.Address()
	wallet.address = &address
	wallet.mnemonic = &mnemonic
}

func (wallet *VsysWallet) GetMnemonic() string {
	if wallet.mnemonic == nil {
		return ""
	}
	return *wallet.mnemonic
}

func (wallet *VsysWallet) BuildFromAddress(address string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.address = &address
}

func (wallet *VsysWallet) GetAddress() string {
	address := wallet.address
	if address == nil {
		return ""
	}
	return *address
}
