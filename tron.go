package go_wallet

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/andnux/go-wallet/base58"
	"github.com/btcsuite/btcd/btcec"
	"github.com/eoscanada/eos-go/ecc"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/crypto/sha3"
	"strconv"
	"strings"
)

type TronWallet struct {
	Test bool
}

func (wallet *TronWallet) Name() string {
	return "Tron"
}

func (wallet *TronWallet) Signature(data []byte, privateKey string) []byte {
	key, err := ecc.NewPrivateKey(privateKey)
	if err != nil {
		fmt.Println(err)
	}
	out, err := key.Sign(data)
	if err != nil {
		fmt.Println(err)
	}
	return out.Content
}

func (wallet *TronWallet) Generate() string {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return wallet.GenerateByMnemonic(mnemonic, "m/44'/195'/0'/0/0")
}
func (wallet *TronWallet) GenerateByPrivateKey(privateKey string) string {
	if !strings.HasPrefix(privateKey, "0x") {
		privateKey = "0x" + privateKey
	}
	keybyte, err := hexutil.Decode(privateKey)
	if err != nil {
		panic(err)
	}
	puk, pub := btcec.PrivKeyFromBytes(btcec.S256(), keybyte)
	serialize := pub.SerializeUncompressed()
	acc := Account{PrivateKey: "0x" + hex.EncodeToString(puk.Serialize()),
		PublicKey: "0x" + hex.EncodeToString(serialize),
		Address:   getAddress(serialize)}
	b, err := json.Marshal(acc)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}

func getAddress(puk []byte) string {
	hash := sha3.NewLegacyKeccak256()
	data := hash.Sum(puk)
	combine := bytesCombine([]byte{byte(41)}, data)
	init := combine[len(combine)-20 : len(combine)]
	sum := hash.Sum(hash.Sum(init))[0:4]
	data = base58.Encode(bytesCombine(init, sum))
	return string(bytesCombine([]byte{'T'}, data))
}

func (wallet *TronWallet) GenerateByMnemonic(mnemonic string, path string) string {
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
	key, err := bip44.NewKeyFromMnemonic(mnemonic, 0x800000c3, uint32(account), uint32(chain), uint32(address))
	if err != nil {
		fmt.Println(err)
	}
	puk := key.PublicKey().Key
	acc := Account{PrivateKey: "0x" + hex.EncodeToString(key.Key),
		PublicKey: "0x" + hex.EncodeToString(puk),
		Address:   getAddress(puk),
		Mnemonic:  mnemonic}
	b, err := json.Marshal(acc)
	if err != nil {
		fmt.Println(err)
	}
	return string(b)
}
