package vsys

import (
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/andnux/go-wallet"
	"github.com/walkbean/vsys-sdk-go/vsys"
)

type VsysWallet struct {
	Test       bool
	privateKey *string
	publicKey  *string
	mnemonic   *string
	address    *string
}

func (wallet *VsysWallet) Sign(data []byte) (signed []byte) {
	if wallet.privateKey == nil {
		return signed
	}
	var account *vsys.Account
	if wallet.Test {
		account = vsys.InitAccount(vsys.Testnet)
	} else {
		account = vsys.InitAccount(vsys.Mainnet)
	}
	account.BuildFromPrivateKey(*wallet.privateKey)
	return []byte(account.SignData(data))
}

func (wallet *VsysWallet) FromGenerate() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.FromMnemonic(mnemonic)
}

func (wallet *VsysWallet) FromPrivateKey(privateKey string) {
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
func (wallet *VsysWallet) FromPublicKey(publicKey string) {
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

func (wallet *VsysWallet) FromMnemonic(mnemonic string) {
	wallet.FromMnemonicAndPathAndNonce(mnemonic, "m/44'/360'/0'/0/0", 0)
}

func (wallet *VsysWallet) FromMnemonicAndPath(mnemonic string, path string) {
	wallet.FromMnemonicAndPathAndNonce(mnemonic, path, 0)
}

func (wallet *VsysWallet) FromMnemonicAndPathAndNonce(mnemonic string, path string, nonce int) {
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

func (wallet *VsysWallet) FromAddress(address string) {
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
