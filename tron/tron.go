package tron

import (
	"encoding/hex"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/andnux/go-wallet"
	"github.com/btcsuite/btcutil/base58"
	"github.com/ethereum/go-ethereum/crypto"
	tron "github.com/sasaxie/go-client-api/common/crypto"
	"golang.org/x/crypto/sha3"
)

type TronWallet struct {
	Test       bool
	privateKey *string
	publicKey  *string
	mnemonic   *string
	keystore   *string
	address    *string
}

func (wallet *TronWallet) Sign(data []byte) (signed []byte) {
	privateKey := wallet.privateKey
	if privateKey == nil {
		return signed
	}
	pk, err := crypto.HexToECDSA(*privateKey)
	if err != nil {
		fmt.Println(err)
	}
	sig, err := crypto.Sign(data, pk)
	if err != nil {
		fmt.Println(err)
	}
	return sig
}

func (wallet *TronWallet) BuildFromRandomGenerate() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.BuildFromMnemonic(mnemonic)
}

func (wallet *TronWallet) GetPrivateKey() string {
	key := wallet.privateKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *TronWallet) BuildFromPrivateKey(privateKey string) {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println(err)
	}
	wallet.privateKey = &privateKey
	bytes := crypto.CompressPubkey(&key.PublicKey)
	hexPubKey := hex.EncodeToString(bytes)
	wallet.publicKey = &hexPubKey
	address := hex.EncodeToString(tron.PubkeyToAddress(key.PublicKey).Bytes())
	wallet.address = &address
}

func (wallet *TronWallet) BuildFromPublicKey(publicKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	puk, err := hex.DecodeString(publicKey)
	if err != nil {
		fmt.Println(err)
	}
	key, err := crypto.DecompressPubkey(puk)
	if err != nil {
		fmt.Println(err)
	}
	wallet.publicKey = &publicKey
	address := hex.EncodeToString(tron.PubkeyToAddress(*key).Bytes())
	wallet.address = &address
}

func (wallet *TronWallet) GetPublicKey() string {
	key := wallet.publicKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *TronWallet) BuildFromMnemonic(mnemonic string) {
	wallet.BuildFromMnemonicAndPath(mnemonic, "m/44'/195'/0'/0/0")
}

func (wallet *TronWallet) BuildFromMnemonicAndPath(mnemonic string, path string) {
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
	address := hex.EncodeToString(tron.PubkeyToAddress(privateKey.PublicKey).Bytes())
	wallet.address = &address
}

func (wallet *TronWallet) GetMnemonic() string {
	if wallet.mnemonic == nil {
		return ""
	}
	return *wallet.mnemonic
}

func (wallet *TronWallet) BuildFromAddress(address string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.address = &address
}

func (wallet *TronWallet) GetAddress() string {
	if wallet.address == nil {
		return ""
	}
	return *wallet.address
}

func (wallet *TronWallet) ConvertToTAddress() string {
	if wallet.address == nil {
		return ""
	}
	bytes, err := hex.DecodeString(*wallet.address)
	if err != nil {
		return ""
	}
	var sum []byte
	hash := sha3.New256()
	sum = hash.Sum(bytes)
	digest := sha3.Sum256(sum)
	data := go_wallet.BytesCombine(bytes, digest[0:4])
	return base58.Encode(data)
}

func (wallet *TronWallet) ConvertTo4Address() string {
	if wallet.address == nil {
		return ""
	}
	bytes := base58.Decode(*wallet.address)
	if len(bytes) < 4 {
		return ""
	}
	var row = bytes[0 : len(bytes)-4]
	var sum []byte
	hash := sha3.New256()
	sum = hash.Sum(row)
	digest := sha3.Sum256(sum)
	if digest[0] == bytes[len(row)+0] &&
		digest[1] == bytes[len(row)+1] &&
		digest[2] == bytes[len(row)+2] &&
		digest[3] == bytes[len(row)+3] {
		return hex.EncodeToString(row)
	}
	return ""
}
