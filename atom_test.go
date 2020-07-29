package go_wallet

import (
	"encoding/hex"
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestAtomWallet_Sign(t *testing.T) {
	wallet := AtomWallet{}
	key := "aff9378bdbf6aed3bea25780fa97c418d90164148515fc88b154b50c476dd8bd"
	wallet.BuildFromPrivateKey(key)
	signed := wallet.Sign([]byte("123123"))
	s := hex.EncodeToString(signed)
	fmt.Println(s)
}

func TestAtomWallet_BuildFromPublicKey(t *testing.T) {
	wallet := AtomWallet{}
	hexByte := "eb5ae9872102b359503f7d860242a87d9af308de49854f65c8dd85e12a34044e24357f16f597"
	wallet.BuildFromPublicKey(hexByte)
	assert.Equal(t, "cosmos1pysu03u65wjvrn77dvkwfzt94ew4eem893ye57", wallet.GetAddress())
}

func TestAtomWallet_BuildFromRandomGenerate(t *testing.T) {
	wallet := AtomWallet{}
	wallet.BuildFromRandomGenerate()
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//wave depend card please wage identify obey stadium hub blast engine win
	//aff9378bdbf6aed3bea25780fa97c418d90164148515fc88b154b50c476dd8bd
	//eb5ae9872102b359503f7d860242a87d9af308de49854f65c8dd85e12a34044e24357f16f597
	//cosmos1pysu03u65wjvrn77dvkwfzt94ew4eem893ye57
}

func TestAtomWallet_BuildFromMnemonic(t *testing.T) {
	wallet := AtomWallet{}
	wallet.BuildFromMnemonic("wave depend card please wage identify obey stadium hub blast engine win")
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "aff9378bdbf6aed3bea25780fa97c418d90164148515fc88b154b50c476dd8bd", wallet.GetPrivateKey())
	assert.Equal(t, "eb5ae9872102b359503f7d860242a87d9af308de49854f65c8dd85e12a34044e24357f16f597", wallet.GetPublicKey())
	assert.Equal(t, "cosmos1pysu03u65wjvrn77dvkwfzt94ew4eem893ye57", wallet.GetAddress())
}

func TestAtomWallet_BuildFromPrivateKey(t *testing.T) {
	wallet := AtomWallet{}
	wallet.BuildFromPrivateKey("aff9378bdbf6aed3bea25780fa97c418d90164148515fc88b154b50c476dd8bd")
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "eb5ae9872102b359503f7d860242a87d9af308de49854f65c8dd85e12a34044e24357f16f597", wallet.GetPublicKey())
	assert.Equal(t, "cosmos1pysu03u65wjvrn77dvkwfzt94ew4eem893ye57", wallet.GetAddress())
}

func TestTemplateWallet_BuildFromMnemonicAndPath(t *testing.T) {
	wallet := AtomWallet{}
	mnemonic := "wave depend card please wage identify obey stadium hub blast engine win"
	wallet.BuildFromMnemonicAndPath(mnemonic, "m/44'/118'/0'/0/0")
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "aff9378bdbf6aed3bea25780fa97c418d90164148515fc88b154b50c476dd8bd", wallet.GetPrivateKey())
	assert.Equal(t, "eb5ae9872102b359503f7d860242a87d9af308de49854f65c8dd85e12a34044e24357f16f597", wallet.GetPublicKey())
	assert.Equal(t, "cosmos1pysu03u65wjvrn77dvkwfzt94ew4eem893ye57", wallet.GetAddress())
}
