package ltc

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/andnux/go-wallet"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ltcsuite/ltcd/btcec"
	"github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcutil"
)

type LtcWallet struct {
	Test       bool
	privateKey *string
	publicKey  *string
	mnemonic   *string
	address    *string
}

func (wallet *LtcWallet) Sign(data []byte) (signed []byte) {
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

func (wallet *LtcWallet) FromGenerate() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.FromMnemonic(mnemonic)
}

func (wallet *LtcWallet) FromPrivateKey(privateKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	params := wallet.getParams()
	wif, err := ltcutil.DecodeWIF(privateKey)
	if err != nil {
		fmt.Println(err)
	}
	pubKey := wif.SerializePubKey()
	hash160 := ltcutil.Hash160(pubKey)
	addressPubKey, err := ltcutil.NewAddressPubKeyHash(hash160, &params)
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

func (wallet *LtcWallet) GetPrivateKey() string {
	key := wallet.privateKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *LtcWallet) FromPublicKey(publicKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.publicKey = &publicKey
	params := wallet.getParams()
	pubKey, err := hex.DecodeString(publicKey)
	if err != nil {
		fmt.Println(err)
	}
	hash160 := ltcutil.Hash160(pubKey)
	addressPubKey, err := ltcutil.NewAddressPubKeyHash(hash160, &params)
	if err != nil {
		fmt.Println(err)
	}
	a := addressPubKey.EncodeAddress()
	wallet.address = &a
}

func (wallet *LtcWallet) GetPublicKey() string {
	key := wallet.publicKey
	if key == nil {
		return ""
	}
	return *key
}
func (wallet *LtcWallet) FromMnemonic(mnemonic string) {
	wallet.FromMnemonicAndPath(mnemonic, "m/44'/2'/0'/0/0")
}

func (wallet *LtcWallet) getParams() chaincfg.Params {
	if wallet.Test {
		return chaincfg.TestNet4Params
	} else {
		return chaincfg.MainNetParams
	}
}

func (wallet *LtcWallet) FromMnemonicAndPath(mnemonic string, path string) {
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
	params := wallet.getParams()
	wif, err := ltcutil.NewWIF(puk, &params, false)
	if err != nil {
		fmt.Println(err)
	}
	s := wif.String()
	wallet.privateKey = &s
	wallet.mnemonic = &mnemonic
	hash160 := ltcutil.Hash160(publicKey)
	addressPubKey, err := ltcutil.NewAddressPubKeyHash(hash160, &params)
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

func (wallet *LtcWallet) GetMnemonic() string {
	if wallet.mnemonic == nil {
		return ""
	}
	return *wallet.mnemonic
}

func (wallet *LtcWallet) FromAddress(address string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.address = &address
}

func (wallet *LtcWallet) GetAddress() string {
	if wallet.address == nil {
		return ""
	}
	return *wallet.address
}
