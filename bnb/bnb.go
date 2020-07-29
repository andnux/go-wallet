package bnb

import (
	"encoding/hex"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/andnux/go-wallet"
	"github.com/binance-chain/go-sdk/common/types"
	"github.com/binance-chain/go-sdk/keys"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

type BnbWallet struct {
	Test       bool
	privateKey *string
	publicKey  *string
	mnemonic   *string
	address    *string
}

func (wallet *BnbWallet) Sign(data []byte) (signed []byte) {
	if wallet.privateKey == nil {
		return signed
	}
	manager, err := keys.NewPrivateKeyManager(*wallet.privateKey)
	if err != nil {
		fmt.Println(err)
	}
	signed, err = manager.GetPrivKey().Sign(data)
	if err != nil {
		fmt.Println(err)
	}
	return signed
}

func (wallet *BnbWallet) BuildFromRandomGenerate() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.BuildFromMnemonic(mnemonic)
}

func (wallet *BnbWallet) BuildFromPrivateKey(privateKey string) {
	wallet.publicKey = nil
	wallet.privateKey = &privateKey
	wallet.mnemonic = nil
	manager, err := keys.NewPrivateKeyManager(*wallet.privateKey)
	if err != nil {
		fmt.Println(err)
	}
	addr := manager.GetAddr().String()
	wallet.address = &addr
}

func (wallet *BnbWallet) GetPrivateKey() string {
	key := wallet.privateKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *BnbWallet) BuildFromPublicKey(publicKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.address = nil
	wallet.mnemonic = nil
	wallet.publicKey = &publicKey
	pubKey, err := hex.DecodeString(*wallet.publicKey)
	if err != nil {
		fmt.Println(err)
	}
	var ptr secp256k1.PubKeySecp256k1
	err = go_wallet.MainCodec.UnmarshalBinaryBare(pubKey, &ptr)
	if err != nil {
		panic(err)
	}
	addr := types.AccAddress(ptr.Address()).String()
	wallet.address = &addr
}

func (wallet *BnbWallet) GetPublicKey() string {
	key := wallet.publicKey
	if key == nil {
		return ""
	}
	return *key
}
func (wallet *BnbWallet) BuildFromMnemonic(mnemonic string) {
	wallet.BuildFromMnemonicAndPath(mnemonic, "m/44'/714'/0'/0/0")
}

func (wallet *BnbWallet) BuildFromMnemonicAndPath(mnemonic string, path string) {
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
	hexKey := hex.EncodeToString(key.Key)
	wallet.privateKey = &hexKey
	if err != nil {
		panic(err)
	}
	var keyBytesArray [32]byte
	copy(keyBytesArray[:], key.Key[:32])
	priKey := secp256k1.PrivKeySecp256k1(keyBytesArray)
	tmpPubKey := priKey.PubKey()
	bytes := tmpPubKey.Bytes()
	pubKey := hex.EncodeToString(bytes)
	wallet.publicKey = &pubKey
	addr := types.AccAddress(tmpPubKey.Address()).String()
	wallet.address = &addr
	wallet.mnemonic = &mnemonic
}

func (wallet *BnbWallet) GetMnemonic() string {
	if wallet.mnemonic == nil {
		return ""
	}
	return *wallet.mnemonic
}

func (wallet *BnbWallet) BuildFromAddress(address string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.address = &address
}

func (wallet *BnbWallet) GetAddress() string {
	if wallet.address == nil {
		return ""
	}
	return *wallet.address
}
