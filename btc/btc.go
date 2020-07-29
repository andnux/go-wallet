package btc

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/andnux/go-wallet"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

type BtcWallet struct {
	Test       bool
	privateKey *string
	publicKey  *string
	mnemonic   *string
	address    *string
}

func (wallet *BtcWallet) Sign(data []byte) (signed []byte) {
	if wallet.privateKey == nil {
		return signed
	}
	key, err := hexutil.Decode(*wallet.privateKey)
	if err != nil {
		fmt.Println(err)
	}
	puk, _ := btcec.PrivKeyFromBytes(btcec.S256(), key)
	signature, err := puk.Sign(data)
	if err != nil {
		fmt.Println(err)
	}
	return signature.Serialize()
}

func (wallet *BtcWallet) FromGenerate() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.FromMnemonic(mnemonic)
}

func (wallet *BtcWallet) FromPrivateKey(privateKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	var params chaincfg.Params
	if wallet.Test {
		params = chaincfg.TestNet3Params
	} else {
		params = chaincfg.MainNetParams
	}
	wif, err := btcutil.DecodeWIF(privateKey)
	if err != nil {
		fmt.Println(err)
	}
	pubKey := wif.SerializePubKey()
	hash160 := btcutil.Hash160(pubKey)
	addressPubKey, err := btcutil.NewAddressPubKeyHash(hash160, &params)
	if err != nil {
		fmt.Println(err)
	}
	a := addressPubKey.EncodeAddress()
	hexPubKey := hex.EncodeToString(pubKey)
	wallet.publicKey = &hexPubKey
	priKey := wif.String()
	wallet.privateKey = &priKey
	wallet.address = &a
}

func (wallet *BtcWallet) GetPrivateKey() string {
	key := wallet.privateKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *BtcWallet) FromPublicKey(publicKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.publicKey = &publicKey
	var params chaincfg.Params
	if wallet.Test {
		params = chaincfg.TestNet3Params
	} else {
		params = chaincfg.MainNetParams
	}
	pubKey, err := hex.DecodeString(publicKey)
	if err != nil {
		fmt.Println(err)
	}
	hash160 := btcutil.Hash160(pubKey)
	addressPubKey, err := btcutil.NewAddressPubKeyHash(hash160, &params)
	if err != nil {
		fmt.Println(err)
	}
	a := addressPubKey.EncodeAddress()
	wallet.address = &a
}

func (wallet *BtcWallet) GetPublicKey() string {
	key := wallet.publicKey
	if key == nil {
		return ""
	}
	return *key
}
func (wallet *BtcWallet) FromMnemonic(mnemonic string) {
	wallet.FromMnemonicAndPath(mnemonic, "m/44'/0'/0'/0/0")
}

func (wallet *BtcWallet) FromMnemonicAndPath(mnemonic string, path string) {
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
	puk, pub := btcec.PrivKeyFromBytes(btcec.S256(), key.Key)
	publicKey := pub.SerializeUncompressed()
	pubHex := hex.EncodeToString(publicKey)
	wallet.publicKey = &pubHex
	var params chaincfg.Params
	if wallet.Test {
		params = chaincfg.TestNet3Params
	} else {
		params = chaincfg.MainNetParams
	}
	wif, err := btcutil.NewWIF(puk, &params, false)
	if err != nil {
		fmt.Println(err)
	}
	s := wif.String()
	wallet.privateKey = &s
	wallet.mnemonic = &mnemonic
	hash160 := btcutil.Hash160(publicKey)
	addressPubKey, err := btcutil.NewAddressPubKeyHash(hash160, &params)
	if err != nil {
		panic(errors.New(err.Error()))
	}
	a := addressPubKey.EncodeAddress()
	hexPubKey := hex.EncodeToString(publicKey)
	wallet.publicKey = &hexPubKey
	priKey := wif.String()
	wallet.privateKey = &priKey
	wallet.address = &a
}

func (wallet *BtcWallet) GetMnemonic() string {
	if wallet.mnemonic == nil {
		return ""
	}
	return *wallet.mnemonic
}

func (wallet *BtcWallet) FromAddress(address string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.address = &address
}

func (wallet *BtcWallet) GetAddress() string {
	if wallet.address == nil {
		return ""
	}
	return *wallet.address
}
