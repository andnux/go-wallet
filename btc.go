package go_wallet

import (
	"encoding/json"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/eoscanada/eos-go/btcsuite/btcd/btcec"
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
	keybyte, err := hexutil.Decode(privateKey)
	if err != nil {
		fmt.Println(err)
	}
	puk, pub := btcec.PrivKeyFromBytes(btcec.S256(), keybyte)
	hash160 := btcutil.Hash160(pub.SerializeUncompressed())
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
	publicKey := hexutil.Encode(pub.SerializeUncompressed())
	acc := Account{PrivateKey: hexutil.Encode(puk.Serialize()),
		PublicKey: publicKey,
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
	hash160 := btcutil.Hash160(pub.SerializeUncompressed())
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
	publicKey := hexutil.Encode(pub.SerializeUncompressed())
	acc := Account{PrivateKey: hexutil.Encode(puk.Serialize()),
		PublicKey: publicKey,
		Mnemonic:  mnemonic,
		Address:   a}
	bytes, err := json.Marshal(acc)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}
