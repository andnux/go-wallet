package go_wallet

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"github.com/gcash/bchutil/bech32"
	"golang.org/x/crypto/ripemd160"
)

type AtomWallet struct {
	Test       bool
	privateKey *string
	publicKey  *string
	mnemonic   *string
	address    *string
}

func (wallet *AtomWallet) Sign(data []byte) (signed []byte) {
	if wallet.privateKey == nil {
		return signed
	}
	bytes, err := hex.DecodeString(*wallet.privateKey)
	if err != nil {
		panic(err)
	}
	sum := sha256.Sum256(data)
	hash := []byte{}
	for i, it := range sum {
		hash[i] = it
	}
	sign, err := secp256k1.Sign(hash, bytes)
	if err != nil {
		panic(err)
	}
	return sign
}

func (wallet *AtomWallet) BuildFromRandomGenerate() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.BuildFromMnemonic(mnemonic)
}

func (wallet *AtomWallet) BuildFromPrivateKey(hexKey string) {
	wallet.publicKey = nil
	wallet.privateKey = &hexKey
	wallet.mnemonic = nil
	bytes, err := hex.DecodeString(*wallet.privateKey)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes)
}

func (wallet *AtomWallet) GetPrivateKey() string {
	key := wallet.privateKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *AtomWallet) publicKeyToAddress(hexPublicKey string) string {
	bytes, err := hex.DecodeString(hexPublicKey)
	if err != nil {
		panic(err)
	}
	sum := sha256.Sum256(bytes)
	hash := make([]byte, 32)
	for i, it := range sum {
		hash[i] = it
	}
	h := ripemd160.New()
	hash = h.Sum(hash)
	addr, err := bech32.Encode("cosmos", hash)
	if err != nil {
		panic(err)
	}
	wallet.address = &addr
	return addr
}

func (wallet *AtomWallet) BuildFromPublicKey(publicKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.publicKey = &publicKey
	addr := wallet.publicKeyToAddress(publicKey)
	wallet.address = &addr
}

func (wallet *AtomWallet) GetPublicKey() string {
	key := wallet.publicKey
	if key == nil {
		return ""
	}
	return *key
}
func (wallet *AtomWallet) BuildFromMnemonic(mnemonic string) {
	wallet.BuildFromMnemonicAndPath(mnemonic, "m/44'/118'/0'/0/0")
}

func (wallet *AtomWallet) BuildFromMnemonicAndPath(mnemonic string, path string) {
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
	privateKey := hex.EncodeToString(key.Key)
	wallet.privateKey = &privateKey
	publicKey := hex.EncodeToString(key.PublicKey().Key)
	wallet.publicKey = &publicKey
	addr := wallet.publicKeyToAddress(publicKey)
	wallet.address = &addr
}

func (wallet *AtomWallet) GetMnemonic() string {
	if wallet.mnemonic == nil {
		return ""
	}
	return *wallet.mnemonic
}

func (wallet *AtomWallet) BuildFromAddress(address string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.address = &address
}

func (wallet *AtomWallet) GetAddress() string {
	if wallet.address == nil {
		return ""
	}
	return *wallet.address
}
