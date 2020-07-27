package go_wallet

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/andnux/go-wallet/filcoin/address"
	"github.com/andnux/go-wallet/filcoin/crypto"
)

type FileCoinWallet struct {
	Test       bool
	privateKey *string
	publicKey  *string
	mnemonic   *string
	address    *string
}

func (wallet *FileCoinWallet) Sign(data []byte) (signed []byte, err error) {
	privateKey := wallet.privateKey
	if privateKey == nil {
		return signed, errors.New("请先导入私钥")
	}
	keyHex, err := hex.DecodeString(*privateKey)
	if err != nil {
		fmt.Println(err)
	}
	signed, err = crypto.Sign(keyHex, data)
	if err != nil {
		fmt.Println(err)
	}
	return signed, nil
}

func (wallet *FileCoinWallet) BuildFromRandomGenerate() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.BuildFromMnemonic(mnemonic)
}

func (wallet *FileCoinWallet) BuildFromPrivateKey(privateKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
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
	wallet.privateKey = &privateKey
	sss := hex.EncodeToString(publicKey)
	wallet.publicKey = &sss
	addr := k1Address.String(network)
	wallet.address = &addr
}

func (wallet *FileCoinWallet) GetPrivateKey() string {
	key := wallet.privateKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *FileCoinWallet) BuildFromPublicKey(publicKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	bytes, err := hex.DecodeString(publicKey)
	k1Address, err := address.NewSecp256k1Address(bytes)
	if err != nil {
		fmt.Println(err)
	}
	var network address.Network
	if wallet.Test {
		network = address.Testnet
	} else {
		network = address.Mainnet
	}
	wallet.publicKey = &publicKey
	addr := k1Address.String(network)
	wallet.address = &addr
}

func (wallet *FileCoinWallet) GetPublicKey() string {
	key := wallet.publicKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *FileCoinWallet) GetMnemonic() string {
	if wallet.mnemonic == nil {
		return ""
	}
	return *wallet.mnemonic
}

func (wallet *FileCoinWallet) BuildFromMnemonic(mnemonic string) {
	wallet.BuildFromMnemonicAndPath(mnemonic, "m/44'/461'/0'/0/0")
}

func (wallet *FileCoinWallet) BuildFromMnemonicAndPath(mnemonic string, path string) {
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
	publicKey := crypto.PublicKey(key.Key)
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
	privateKey := hex.EncodeToString(key.Key)
	wallet.privateKey = &privateKey
	hexPuk := hex.EncodeToString(publicKey)
	wallet.publicKey = &hexPuk
	addr := k1Address.String(network)
	wallet.address = &addr
	wallet.mnemonic = &mnemonic
}

func (wallet *FileCoinWallet) BuildFromAddress(address string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.address = &address
}

func (wallet *FileCoinWallet) GetAddress() string {
	if wallet.address == nil {
		return ""
	}
	return *wallet.address
}
