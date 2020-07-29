package xrp

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/andnux/go-wallet"
	"github.com/mr-tron/base58"
	"golang.org/x/crypto/ripemd160"
)

type XrpWallet struct {
	Test       bool
	privateKey *string
	publicKey  *string
	mnemonic   *string
	address    *string
}

func (wallet *XrpWallet) Sign(data []byte) (signed []byte) {
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

func (wallet *XrpWallet) BuildFromRandomGenerate() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.BuildFromMnemonic(mnemonic)
}

func (wallet *XrpWallet) BuildFromPrivateKey(hexKey string) {
	wallet.publicKey = nil
	wallet.privateKey = &hexKey
	wallet.mnemonic = nil
	bytes, err := hex.DecodeString(*wallet.privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes)
}

func (wallet *XrpWallet) GetPrivateKey() string {
	key := wallet.privateKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *XrpWallet) publicKeyToAddress(hexPublicKey string) string {
	bytes, err := hex.DecodeString(hexPublicKey)
	if err != nil {
		panic(err)
	}
	if len(bytes) != 33 {
		panic(errors.New("Public Key Error"))
	}
	hash := sha256.New()
	hash.Write(bytes)
	sha256Hash := hash.Sum(nil)
	h := ripemd160.New()
	h.Write(sha256Hash)
	ripemdHash := h.Sum(nil)
	var payload = make([]byte, 0)
	payload = append(payload, byte(0x00))
	payload = append(payload, ripemdHash...)
	hash.Reset()
	hash.Write(payload)
	hash1 := hash.Sum(nil)
	hash.Reset()
	hash.Write(hash1)
	checksum := hash.Sum(nil)[0:4]
	payload = append(payload, checksum...)
	addr := base58.Encode(payload)
	wallet.address = &addr
	return addr
}

func (wallet *XrpWallet) BuildFromPublicKey(publicKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.publicKey = &publicKey
	addr := wallet.publicKeyToAddress(publicKey)
	wallet.address = &addr
}

func (wallet *XrpWallet) GetPublicKey() string {
	key := wallet.publicKey
	if key == nil {
		return ""
	}
	return *key
}
func (wallet *XrpWallet) BuildFromMnemonic(mnemonic string) {
	wallet.BuildFromMnemonicAndPath(mnemonic, "m/44'/145'/0'/0/0")
}

func (wallet *XrpWallet) BuildFromMnemonicAndPath(mnemonic string, path string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	parser, err := go_wallet.Bip44Parser(path)
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

func (wallet *XrpWallet) GetMnemonic() string {
	if wallet.mnemonic == nil {
		return ""
	}
	return *wallet.mnemonic
}

func (wallet *XrpWallet) BuildFromAddress(address string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.address = &address
}

func (wallet *XrpWallet) GetAddress() string {
	if wallet.address == nil {
		return ""
	}
	return *wallet.address
}
