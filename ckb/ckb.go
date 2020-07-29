package ckb

import (
	"encoding/hex"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/andnux/go-wallet"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/blake2b"
	"github.com/nervosnetwork/ckb-sdk-go/crypto/secp256k1"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
)

type CkbWallet struct {
	Test       bool
	privateKey *string
	publicKey  *string
	mnemonic   *string
	address    *string
}

func (wallet *CkbWallet) Sign(data []byte) (signed []byte) {
	if wallet.privateKey == nil {
		return signed
	}
	key, err := secp256k1.HexToKey(*wallet.privateKey)
	if err != nil {
		fmt.Println(err)
	}
	sign, err := key.Sign(data)
	if err != nil {
		fmt.Println(err)
	}
	return sign
}

func (wallet *CkbWallet) FromGenerate() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.FromMnemonic(mnemonic)
}

func (wallet *CkbWallet) FromPrivateKey(privateKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	key, err := secp256k1.HexToKey(privateKey)
	if err != nil {
		fmt.Println(err)
	}
	pubKey := key.PubKey()
	var params address.Mode
	if wallet.Test {
		params = address.Testnet
	} else {
		params = address.Mainnet
	}
	hexPubKey := hex.EncodeToString(pubKey)
	bytes, err := blake2b.Blake160(pubKey)
	script := &types.Script{
		CodeHash: types.HexToHash(transaction.SECP256K1_BLAKE160_SIGHASH_ALL_TYPE_HASH),
		HashType: types.HashTypeType,
		Args:     bytes,
	}
	addr, err := address.Generate(params, script)
	if err != nil {
		fmt.Println(err)
	}
	wallet.publicKey = &hexPubKey
	priKey := hex.EncodeToString(key.Bytes())
	wallet.privateKey = &priKey
	wallet.address = &addr
}

func (wallet *CkbWallet) GetPrivateKey() string {
	key := wallet.privateKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *CkbWallet) FromPublicKey(publicKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.publicKey = &publicKey
	pubKey, err := hex.DecodeString(publicKey)
	if err != nil {
		fmt.Println(err)
	}
	var params address.Mode
	if wallet.Test {
		params = address.Testnet
	} else {
		params = address.Mainnet
	}
	bytes, err := blake2b.Blake160(pubKey)
	script := &types.Script{
		CodeHash: types.HexToHash(transaction.SECP256K1_BLAKE160_SIGHASH_ALL_TYPE_HASH),
		HashType: types.HashTypeType,
		Args:     bytes,
	}
	addr, err := address.Generate(params, script)
	if err != nil {
		fmt.Println(err)
	}
	wallet.address = &addr
}

func (wallet *CkbWallet) GetPublicKey() string {
	key := wallet.publicKey
	if key == nil {
		return ""
	}
	return *key
}
func (wallet *CkbWallet) FromMnemonic(mnemonic string) {
	wallet.FromMnemonicAndPath(mnemonic, "m/44'/309'/0'/0/0")
}

func (wallet *CkbWallet) FromMnemonicAndPath(mnemonic string, path string) {
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
	ckbKey, err := secp256k1.ToKey(key.Key)
	if err != nil {
		fmt.Println(err)
	}
	pubKey := ckbKey.PubKey()
	var params address.Mode
	if wallet.Test {
		params = address.Testnet
	} else {
		params = address.Mainnet
	}
	hexPubKey := hex.EncodeToString(pubKey)
	bytes, err := blake2b.Blake160(pubKey)
	if err != nil {
		fmt.Println(err)
	}
	script := &types.Script{
		CodeHash: types.HexToHash(transaction.SECP256K1_BLAKE160_SIGHASH_ALL_TYPE_HASH),
		HashType: types.HashTypeType,
		Args:     bytes,
	}
	addr, err := address.Generate(params, script)
	if err != nil {
		fmt.Println(err)
	}
	wallet.publicKey = &hexPubKey
	priKey := hex.EncodeToString(ckbKey.Bytes())
	wallet.privateKey = &priKey
	wallet.address = &addr
	wallet.mnemonic = &mnemonic
}

func (wallet *CkbWallet) GetMnemonic() string {
	if wallet.mnemonic == nil {
		return ""
	}
	return *wallet.mnemonic
}

func (wallet *CkbWallet) FromAddress(address string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.address = &address
}

func (wallet *CkbWallet) GetAddress() string {
	if wallet.address == nil {
		return ""
	}
	return *wallet.address
}
