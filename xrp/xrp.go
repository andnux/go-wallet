package xrp

import (
	"encoding/hex"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/andnux/go-wallet"
	"github.com/minio/sha256-simd"
	"github.com/rubblelabs/ripple/crypto"
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
	ecdsaKey, err := crypto.NewECDSAKey(bytes)
	if err != nil {
		panic(err)
	}
	sign, err := ecdsaKey.Sign(data)
	if err != nil {
		panic(err)
	}
	return sign.Serialize()
}

func (wallet *XrpWallet) FromGenerate() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.FromMnemonic(mnemonic)
}

func (wallet *XrpWallet) FromPrivateKey(hexKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	bytes, err := hex.DecodeString(hexKey)
	if err != nil {
		panic(err)
	}

	fmt.Println(hex.EncodeToString(bytes))
	//TODO 私钥转换有问题
	//ecdsaKey, err := crypto.NewECDSAKey(bytes)
	//if err != nil {
	//	panic(err)
	//}
	//key := hex.EncodeToString(ecdsaKey.PrivateKey.Serialize())
	//wallet.privateKey = &key
	//pukWaw := ecdsaKey.PubKey().SerializeCompressed()
	//pubicKey := hex.EncodeToString(pukWaw)
	//wallet.publicKey = &pubicKey
	//addr := wallet.publicKeyToAddress(pubicKey)
	//wallet.address = &addr
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
	h := sha256.New()
	h.Write(bytes)
	sum := h.Sum(nil)
	h2 := ripemd160.New()
	h2.Write(sum)
	res := h2.Sum(nil)
	hash, err := crypto.NewAccountId(res)
	if err != nil {
		panic(err)
	}
	addr := hash.String()
	wallet.address = &addr
	return addr
}

func (wallet *XrpWallet) FromPublicKey(publicKey string) {
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
func (wallet *XrpWallet) FromMnemonic(mnemonic string) {
	wallet.FromMnemonicAndPath(mnemonic, "m/44'/144'/0'/0/0")
}

func (wallet *XrpWallet) FromMnemonicAndPath(mnemonic string, path string) {
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
	privateKey := hex.EncodeToString(key.Key)
	pubicKey := hex.EncodeToString(key.PublicKey().Key)
	wallet.publicKey = &pubicKey
	wallet.privateKey = &privateKey
	addr := wallet.publicKeyToAddress(pubicKey)
	wallet.address = &addr
	wallet.mnemonic = &mnemonic
}

func (wallet *XrpWallet) GetMnemonic() string {
	if wallet.mnemonic == nil {
		return ""
	}
	return *wallet.mnemonic
}

func (wallet *XrpWallet) FromAddress(address string) {
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
