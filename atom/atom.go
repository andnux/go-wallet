package atom

import (
	"encoding/hex"
	"github.com/FactomProject/go-bip39"
	"github.com/FactomProject/go-bip44"
	"github.com/andnux/go-wallet"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/secp256k1"
)

type AtomWallet struct {
	Test       bool
	privateKey *string
	publicKey  *string
	mnemonic   *string
	address    *string
}

func (wallet *AtomWallet) Sign(data []byte) (signed []byte) {
	if wallet.privateKey == nil {
		return signed
	}
	t, err := hex.DecodeString(*wallet.privateKey)
	if err != nil {
		panic(err)
	}
	var keyBytesArray [32]byte
	copy(keyBytesArray[:], t)
	priKey := secp256k1.PrivKeySecp256k1(keyBytesArray)
	signed, err = priKey.Sign(data)
	if err != nil {
		panic(err)
	}
	return signed
}

func (wallet *AtomWallet) FromGenerate() {
	entropy, _ := bip39.NewEntropy(128)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	wallet.FromMnemonic(mnemonic)
}

func (wallet *AtomWallet) FromPrivateKey(hexKey string) {
	wallet.publicKey = nil
	wallet.privateKey = &hexKey
	wallet.mnemonic = nil
	t, err := hex.DecodeString(*wallet.privateKey)
	if err != nil {
		panic(err)
	}
	var keyBytesArray [32]byte
	copy(keyBytesArray[:], t)
	priKey := secp256k1.PrivKeySecp256k1(keyBytesArray)
	tmpPubKey := priKey.PubKey()
	bytes := tmpPubKey.Bytes()
	pubKey := hex.EncodeToString(bytes)
	wallet.publicKey = &pubKey
	addr := wallet.publicKeyToAddress(pubKey)
	wallet.address = &addr
}

func (wallet *AtomWallet) GetPrivateKey() string {
	key := wallet.privateKey
	if key == nil {
		return ""
	}
	return *key
}

func (wallet *AtomWallet) publicKeyToAddress(hexPublicKey string) string {
	bytes, err := hex.DecodeString(hexPublicKey)
	if err != nil {
		panic(err)
	}
	var b [33]byte
	for i := range b {
		b[i] = bytes[i]
	}
	var ptr secp256k1.PubKeySecp256k1
	err = go_wallet.MainCodec.UnmarshalBinaryBare(bytes, &ptr)
	if err != nil {
		panic(err)
	}
	addr := types.AccAddress(ptr.Address()).String()
	wallet.address = &addr
	return addr
}

func (wallet *AtomWallet) FromPublicKey(publicKey string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.publicKey = &publicKey
	addr := wallet.publicKeyToAddress(publicKey)
	wallet.address = &addr
}

func (wallet *AtomWallet) GetPublicKey() string {
	key := wallet.publicKey
	if key == nil {
		return ""
	}
	return *key
}
func (wallet *AtomWallet) FromMnemonic(mnemonic string) {
	wallet.FromMnemonicAndPath(mnemonic, "m/44'/118'/0'/0/0")
}

func (wallet *AtomWallet) FromMnemonicAndPath(mnemonic string, path string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	parser, err := go_wallet.Bip44Parser(path)
	if err != nil {
		panic(err)
	}
	key, err := bip44.NewKeyFromMnemonic(mnemonic,
		parser.CoinType, parser.Account,
		parser.Change, parser.AddressIndex)
	if err != nil {
		panic(err)
	}
	privateKey := hex.EncodeToString(key.Key)
	wallet.privateKey = &privateKey
	var keyBytesArray [32]byte
	copy(keyBytesArray[:], key.Key[:32])
	priKey := secp256k1.PrivKeySecp256k1(keyBytesArray)
	tmpPubKey := priKey.PubKey()
	bytes := tmpPubKey.Bytes()
	pubKey := hex.EncodeToString(bytes)
	wallet.publicKey = &pubKey
	addr := wallet.publicKeyToAddress(pubKey)
	wallet.address = &addr
	wallet.mnemonic = &mnemonic
}

func (wallet *AtomWallet) GetMnemonic() string {
	if wallet.mnemonic == nil {
		return ""
	}
	return *wallet.mnemonic
}

func (wallet *AtomWallet) FromAddress(address string) {
	wallet.publicKey = nil
	wallet.privateKey = nil
	wallet.mnemonic = nil
	wallet.address = &address
}

func (wallet *AtomWallet) GetAddress() string {
	if wallet.address == nil {
		return ""
	}
	return *wallet.address
}
