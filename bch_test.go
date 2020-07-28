package go_wallet

import (
	"encoding/hex"
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestBchWallet_Sign(t *testing.T) {
	wallet := BchWallet{}
	wallet.Test = false
	wallet.BuildFromPrivateKey("123")
	sign := wallet.Sign([]byte{1})
	fmt.Println(hex.EncodeToString(sign))
}

func TestBchWallet_BuildFromRandomGenerate(t *testing.T) {
	wallet := BchWallet{}
	wallet.Test = false
	wallet.BuildFromRandomGenerate()
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	//L5dX2Tyxt7hfK9fpLDg269nPXArQRq9TE6fTtSqYQPE5Nr7imoGA
	//036802d3d4110a4b4cc0116daa486b92341182b9fa6cdd03022c83db42fb6ab9a5
	//road orbit idea endorse margin exit solid injury super earn journey vanish
	//bitcoincash:qq5uxsetphm4h6e0w02s90epmm0htdkykcwwresn63
}

func TestBchWallet_BuildFromPrivateKey(t *testing.T) {
	wallet := BchWallet{}
	wallet.Test = false
	wallet.BuildFromPrivateKey("L5dX2Tyxt7hfK9fpLDg269nPXArQRq9TE6fTtSqYQPE5Nr7imoGA")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "bitcoincash:qq5uxsetphm4h6e0w02s90epmm0htdkykcwwresn63", wallet.GetAddress())
}

func TestBchWallet_BuildFromMnemonic(t *testing.T) {
	wallet := BchWallet{}
	wallet.Test = false
	mnemonic := "road orbit idea endorse margin exit solid injury super earn journey vanish"
	wallet.BuildFromMnemonic(mnemonic)
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "L5dX2Tyxt7hfK9fpLDg269nPXArQRq9TE6fTtSqYQPE5Nr7imoGA", wallet.GetPrivateKey())
	assert.Equal(t, "036802d3d4110a4b4cc0116daa486b92341182b9fa6cdd03022c83db42fb6ab9a5", wallet.GetPublicKey())
	assert.Equal(t, "bitcoincash:qq5uxsetphm4h6e0w02s90epmm0htdkykcwwresn63", wallet.GetAddress())
}

func TestBchWallet_BuildFromMnemonicAndPath(t *testing.T) {
	wallet := BchWallet{}
	wallet.Test = false
	mnemonic := "road orbit idea endorse margin exit solid injury super earn journey vanish"
	wallet.BuildFromMnemonicAndPath(mnemonic, "m/44'/145'/0'/0/0")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "L5dX2Tyxt7hfK9fpLDg269nPXArQRq9TE6fTtSqYQPE5Nr7imoGA", wallet.GetPrivateKey())
	assert.Equal(t, "036802d3d4110a4b4cc0116daa486b92341182b9fa6cdd03022c83db42fb6ab9a5", wallet.GetPublicKey())
	assert.Equal(t, "bitcoincash:qq5uxsetphm4h6e0w02s90epmm0htdkykcwwresn63", wallet.GetAddress())
}

func TestBchWallet_BuildFromPublicKey(t *testing.T) {
	wallet := BchWallet{}
	wallet.Test = false
	wallet.BuildFromPublicKey("036802d3d4110a4b4cc0116daa486b92341182b9fa6cdd03022c83db42fb6ab9a5")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "bitcoincash:qq5uxsetphm4h6e0w02s90epmm0htdkykcwwresn63", wallet.GetAddress())
}
