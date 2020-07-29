package eos

import (
	"fmt"
	"testing"
)

func TestEosWallet_BuildFromRandomGenerate(t *testing.T) {
	wallet := EosWallet{}
	wallet.BuildFromRandomGenerate()
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	//topple defense shell defense firm waste help update glove betray actual tower
	//5JUM8rj8ktgwmsDym6b2Qcj2SpCUdcXd5wGNKskatD7ixjHCmRv
	//EOS7CE42LYf6jYC6RrcWP76C3KRGGyHmmPvQ2Sj2VCJKATYfCyKcK
}

func TestEosWallet_BuildFromPrivateKey(t *testing.T) {
	wallet := EosWallet{}
	privateKey := "5JUM8rj8ktgwmsDym6b2Qcj2SpCUdcXd5wGNKskatD7ixjHCmRv"
	wallet.BuildFromPrivateKey(privateKey)
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	//5JUM8rj8ktgwmsDym6b2Qcj2SpCUdcXd5wGNKskatD7ixjHCmRv
	//EOS7CE42LYf6jYC6RrcWP76C3KRGGyHmmPvQ2Sj2VCJKATYfCyKcK
}

func TestEosWallet_BuildFromPublicKey(t *testing.T) {
	wallet := EosWallet{}
	publicKey := "EOS7CE42LYf6jYC6RrcWP76C3KRGGyHmmPvQ2Sj2VCJKATYfCyKcK"
	wallet.BuildFromPublicKey(publicKey)
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
}

func TestEosWallet_BuildFromMnemonic(t *testing.T) {
	wallet := EosWallet{}
	mnemonic := "topple defense shell defense firm waste help update glove betray actual tower"
	wallet.BuildFromMnemonic(mnemonic)
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	//topple defense shell defense firm waste help update glove betray actual tower
	//5JUM8rj8ktgwmsDym6b2Qcj2SpCUdcXd5wGNKskatD7ixjHCmRv
	//EOS7CE42LYf6jYC6RrcWP76C3KRGGyHmmPvQ2Sj2VCJKATYfCyKcK
}
