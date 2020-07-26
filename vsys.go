package go_wallet

import (
	"encoding/json"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/walkbean/vsys-sdk-go/vsys"
	"strconv"
	"strings"
)

type VsysWallet struct {
	Test bool
}

func (wallet *VsysWallet) Name() string {
	return "VSYS"
}

func (wallet *VsysWallet) Signature(data []byte, privateKey string) []byte {
	var account *vsys.Account
	if wallet.Test {
		account = vsys.InitAccount(vsys.Testnet)
	} else {
		account = vsys.InitAccount(vsys.Mainnet)
	}
	account.BuildFromPrivateKey(privateKey)
	return []byte(account.SignData(data))
}

func (wallet *VsysWallet) Generate(test bool) string {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return wallet.GenerateByMnemonic(mnemonic, "m/44'/360'/0'/0/0", test)
}
func (wallet *VsysWallet) GenerateByPrivateKey(privateKey string, test bool) string {
	var account *vsys.Account
	if test {
		account = vsys.InitAccount(vsys.Testnet)
	} else {
		account = vsys.InitAccount(vsys.Mainnet)
	}
	account.BuildFromPrivateKey(privateKey)
	a := Account{PrivateKey: account.PrivateKey(),
		PublicKey: account.PublicKey(),
		Address:   account.Address()}
	bytes, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}

func (wallet *VsysWallet) GenerateByMnemonic(mnemonic string, path string, test bool) string {
	return wallet.GenerateByMnemonicAndNonce(mnemonic, path, 0, test)
}

func (wallet *VsysWallet) GenerateByMnemonicAndNonce(mnemonic string, path string, nonce int, test bool) string {
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
	key, err := bip44.NewKeyFromMnemonic(mnemonic, 0x80000168, uint32(account), uint32(chain), uint32(address))
	if err != nil {
		fmt.Println(err)
	}
	var vsysAccount *vsys.Account
	if test {
		vsysAccount = vsys.InitAccount(vsys.Testnet)
	} else {
		vsysAccount = vsys.InitAccount(vsys.Mainnet)
	}
	vsysAccount.BuildFromSeed(string(key.Key), nonce)
	a := Account{PrivateKey: vsysAccount.PrivateKey(),
		PublicKey: vsysAccount.PublicKey(),
		Mnemonic:  mnemonic,
		Address:   vsysAccount.Address()}
	bytes, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}
