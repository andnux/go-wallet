package go_wallet

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strconv"
	"strings"
)

type BtcWallet struct {
	Test bool
}

func (wallet *BtcWallet) Name() string {
	return "BTC"
}

func (wallet *BtcWallet) Signature(data []byte, privateKey string) []byte {
	key, err := hexutil.Decode(privateKey)
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

func (wallet *BtcWallet) Generate() string {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return wallet.GenerateByMnemonic(mnemonic, "m/44'/0'/0'/0/0")
}
func (wallet *BtcWallet) GenerateByPrivateKey(privateKey string) string {
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
	acc := Account{PrivateKey: wif.String(),
		PublicKey: hex.EncodeToString(pubKey),
		Address:   a}
	bytes, err := json.Marshal(acc)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}

func (wallet *BtcWallet) GenerateByMnemonic(mnemonic string, path string) string {
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
	keyMnemonic, err := bip44.NewKeyFromMnemonic(mnemonic, 0x80000019, uint32(account), uint32(chain), uint32(address))
	if err != nil {
		fmt.Println(err)
	}
	puk, pub := btcec.PrivKeyFromBytes(btcec.S256(), keyMnemonic.Key)
	publicKey := pub.SerializeUncompressed()
	hash160 := btcutil.Hash160(publicKey)
	var params chaincfg.Params
	if wallet.Test {
		params = chaincfg.TestNet3Params
	} else {
		params = chaincfg.MainNetParams
	}
	addressPubKey, err := btcutil.NewAddressPubKeyHash(hash160, &params)
	if err != nil {
		fmt.Println(err)
	}
	a := addressPubKey.EncodeAddress()
	wif, err := btcutil.NewWIF(puk, &params, false)
	if err != nil {
		fmt.Println(err)
	}
	acc := Account{PrivateKey: wif.String(),
		PublicKey: hex.EncodeToString(publicKey),
		Mnemonic:  mnemonic,
		Address:   a}
	bytes, err := json.Marshal(acc)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}
