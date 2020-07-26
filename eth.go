package go_wallet

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"strconv"
	"strings"
)

type EthWallet struct {
	Test bool
}

func (wallet *EthWallet) Name() string {
	return "Eth"
}

func (wallet *EthWallet) Signature(data []byte, privateKey string) []byte {
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		fmt.Println(err)
	}
	sig, err := crypto.Sign(data, key)
	if err != nil {
		fmt.Println(err)
	}
	return sig
}

func (wallet *EthWallet) Generate() string {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return wallet.GenerateByMnemonic(mnemonic, "m/44'/60'/0'/0/0")
}
func (wallet *EthWallet) GenerateByPrivateKey(keyHex string) string {
	privateKey, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		fmt.Println(err)
	}
	compressPubkey := secp256k1.CompressPubkey(privateKey.PublicKey.X,
		privateKey.PublicKey.Y)
	address := hex.EncodeToString(crypto.PubkeyToAddress(
		privateKey.PublicKey).Bytes())
	a := Account{
		PrivateKey: keyHex,
		PublicKey:  hex.EncodeToString(compressPubkey),
		Address:    "0x" + address}
	bytes, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}

func (wallet *EthWallet) GenerateByMnemonic(mnemonic string, path string) string {
	s := strings.Split(path, "/")
	address, err := strconv.ParseUint(strings.ReplaceAll(s[len(s)-1], "'", ""), 0, 1)
	if err != nil {
		fmt.Println(err)
	}
	chain, err := strconv.ParseUint(strings.ReplaceAll(s[len(s)-2], "'", ""), 0, 1)
	if err != nil {
		fmt.Println(err)
	}
	account, err := strconv.ParseUint(strings.ReplaceAll(s[len(s)-3], "'", ""), 0, 1)
	if err != nil {
		fmt.Println(err)
	}
	keyFromMnemonic, err := bip44.NewKeyFromMnemonic(mnemonic, 0x8000003c, uint32(account), uint32(chain), uint32(address))
	if err != nil {
		fmt.Println(err)
	}
	keyHex := hex.EncodeToString(keyFromMnemonic.Key)
	privateKey, err := crypto.HexToECDSA(keyHex)
	if err != nil {
		fmt.Println(err)
	}
	publicKey := crypto.CompressPubkey(&privateKey.PublicKey)
	a := Account{
		PrivateKey: keyHex,
		PublicKey:  hex.EncodeToString(publicKey),
		Mnemonic:   mnemonic,
		Address: "0x" + hex.EncodeToString(
			crypto.PubkeyToAddress(privateKey.PublicKey).Bytes())}
	bytes, err := json.Marshal(a)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}
