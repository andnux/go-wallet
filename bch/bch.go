package bch

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/andnux/go-wallet"
	"github.com/gcash/bchd/bchec"
	"github.com/gcash/bchd/chaincfg"
	"github.com/gcash/bchutil"
	"github.com/schancel/cashaddr-converter/address"
)

type BchWallet struct {
	Test       bool
	privateKey *string
	publicKey  *string
	mnemonic   *string
	address    *string
}

func (wallet *BchWallet) Sign(data []byte) (signed []byte) {
	if wallet.privateKey == nil {
		return signed
	}
	bytes, err := hex.DecodeString(*wallet.privateKey)
	if err != nil {
		panic(err)
	}
	privateKey, _ := bchec.PrivKeyFromBytes(bchec.S256(), bytes)
	sum := sha256.Sum256(data)
	signed = []byte{}
	for i, it := range sum {
		signed[i] = it
	}
	ecdsa, err := privateKey.SignECDSA(signed)
	if err != nil {
		panic(err)
	}
	return ecdsa.Serialize()
}

func (wallet *BchWallet) FromGenerate() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.FromMnemonic(mnemonic)
}

func (wallet *BchWallet) FromPrivateKey(hexKey string) {
	wallet.publicKey = nil
	wallet.privateKey = &hexKey
	wallet.mnemonic = nil
	bytes, err := hex.DecodeString(*wallet.privateKey)
	if err != nil {
		wif, err := bchutil.DecodeWIF(hexKey)
		if err != nil {
			panic(err)
		}
		bytes = wif.PrivKey.Serialize()
		wallet.privateKey = &hexKey
		key := hex.EncodeToString(wif.SerializePubKey())
		wallet.publicKey = &key
	} else {
		privateKey, publicKey := bchec.PrivKeyFromBytes(bchec.S256(), bytes)
		privateKeyHex := hex.EncodeToString(privateKey.Serialize())
		compressed := publicKey.SerializeUncompressed()
		publicKeyHex := hex.EncodeToString(compressed)
		wallet.privateKey = &privateKeyHex
		wallet.publicKey = &publicKeyHex
	}
	addr := wallet.publicKeyToAddress(*wallet.publicKey)
	wallet.address = &addr
}

func (wallet *BchWallet) GetPrivateKey() string {
	key := wallet.privateKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *BchWallet) getParams() chaincfg.Params {
	if wallet.Test {
		return chaincfg.TestNet3Params
	} else {
		return chaincfg.MainNetParams
	}
}

func (wallet *BchWallet) publicKeyToAddress(hexPublicKey string) string {
	bytes, err := hex.DecodeString(hexPublicKey)
	if err != nil {
		panic(err)
	}
	params := wallet.getParams()
	pubKey, err := bchutil.NewAddressPubKey(bytes, &params)
	if err != nil {
		panic(err)
	}
	addr := pubKey.EncodeAddress()
	//下面代码是转换成新地址bitcoincash:XXXXXXXXXX
	a, err := address.NewFromString(addr)
	if err != nil {
		panic(err)
	}
	cashAddress, err := a.CashAddress()
	if err != nil {
		panic(err)
	}
	addr = cashAddress.String()
	wallet.address = &addr
	return addr
}

func (wallet *BchWallet) FromPublicKey(publicKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.publicKey = &publicKey
	addr := wallet.publicKeyToAddress(publicKey)
	wallet.address = &addr
}

func (wallet *BchWallet) GetPublicKey() string {
	key := wallet.publicKey
	if key == nil {
		return ""
	}
	return *key
}
func (wallet *BchWallet) FromMnemonic(mnemonic string) {
	wallet.FromMnemonicAndPath(mnemonic, "m/44'/145'/0'/0/0")
}

func (wallet *BchWallet) FromMnemonicAndPath(mnemonic string, path string) {
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
	privateKey, _ := bchec.PrivKeyFromBytes(bchec.S256(), key.Key)
	params := wallet.getParams()
	wif, err := bchutil.NewWIF(privateKey, &params, true)
	if err != nil {
		panic(err)
	}
	privateKeyHex := wif.String()
	publicKeyHex := hex.EncodeToString(wif.SerializePubKey())
	wallet.privateKey = &privateKeyHex
	wallet.publicKey = &publicKeyHex
	addr := wallet.publicKeyToAddress(*wallet.publicKey)
	wallet.address = &addr
	wallet.mnemonic = &mnemonic
}

func (wallet *BchWallet) GetMnemonic() string {
	if wallet.mnemonic == nil {
		return ""
	}
	return *wallet.mnemonic
}

func (wallet *BchWallet) FromAddress(address string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.address = &address
}

func (wallet *BchWallet) GetAddress() string {
	if wallet.address == nil {
		return ""
	}
	return *wallet.address
}
