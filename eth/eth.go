package eth

import (
	"encoding/hex"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/andnux/go-wallet"
	"github.com/ethereum/go-ethereum/crypto"
)

type EthWallet struct {
	Test       bool
	privateKey *string
	publicKey  *string
	mnemonic   *string
	keystore   *string
	address    *string
}

func (wallet *EthWallet) Name() string {
	return "Eth"
}

func (wallet *EthWallet) Sign(data []byte) (signed []byte) {
	privateKey := wallet.privateKey
	if privateKey == nil {
		return signed
	}
	key, err := crypto.HexToECDSA(*privateKey)
	if err != nil {
		fmt.Println(err)
	}
	sig, err := crypto.Sign(data, key)
	if err != nil {
		fmt.Println(err)
	}
	return sig
}

func (wallet *EthWallet) FromGenerate() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.FromMnemonicAndPath(mnemonic, "m/44'/60'/0'/0/0")
}

func (wallet *EthWallet) GetPrivateKey() string {
	key := wallet.privateKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *EthWallet) FromPrivateKey(privateKey string) {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println(err)
	}
	publicKey := crypto.CompressPubkey(&pk.PublicKey)
	wallet.privateKey = &privateKey
	puk := hex.EncodeToString(publicKey)
	wallet.publicKey = &puk
	addr := "0x" + hex.EncodeToString(crypto.PubkeyToAddress(pk.PublicKey).Bytes())
	wallet.address = &addr
}

func (wallet *EthWallet) FromPublicKey(publicKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.publicKey = &publicKey
	bytes, err := hex.DecodeString(publicKey)
	if err != nil {
		fmt.Println(err)
	}
	pubkey, err := crypto.DecompressPubkey(bytes)
	if err != nil {
		fmt.Println(err)
	}
	addr := "0x" + hex.EncodeToString(crypto.PubkeyToAddress(*pubkey).Bytes())
	wallet.address = &addr
}

func (wallet *EthWallet) GetPublicKey() string {
	key := wallet.publicKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *EthWallet) FromMnemonic(mnemonic string) {
	wallet.FromMnemonicAndPath(mnemonic, "m/44'/60'/0'/0/0")
}

func (wallet *EthWallet) FromMnemonicAndPath(mnemonic string, path string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.keystore = nil
	wallet.mnemonic = nil
	wallet.address = nil
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
	keyHex := hex.EncodeToString(key.Key)
	privateKey, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		fmt.Println(err)
	}
	publicKey := crypto.CompressPubkey(&privateKey.PublicKey)
	wallet.mnemonic = &mnemonic
	wallet.privateKey = &keyHex
	puk := hex.EncodeToString(publicKey)
	wallet.publicKey = &puk
	addr := "0x" + hex.EncodeToString(crypto.PubkeyToAddress(privateKey.PublicKey).Bytes())
	wallet.address = &addr
}

func (wallet *EthWallet) GetMnemonic() string {
	if wallet.mnemonic == nil {
		return ""
	}
	return *wallet.mnemonic
}

func (wallet *EthWallet) FromAddress(address string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.address = &address
}

func (wallet *EthWallet) GetAddress() string {
	if wallet.address == nil {
		return ""
	}
	return *wallet.address
}
