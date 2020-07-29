package vsys

import (
	"fmt"
	"testing"
)

func TestVsysWallet_FromGenerate(t *testing.T) {
	wallet := VsysWallet{}
	wallet.Test = false
	wallet.FromGenerate()
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//inherit fiction obscure angry scale mouse fine tornado fitness before cricket parade
	//AhSeFuxHydXkDV3Qp1F1onoPTsG6fFeLMsB9K7qBtjwM
	//CXwmA9bYf6AP53jp2jzRG5nCtV8QEPe6AwbbX9a2dM6n
	//AR7EFpUF2bmztKLNPbEL8CT5HpahyiWwwo8
}

func TestVsysWallet_FromPrivateKey(t *testing.T) {
	wallet := VsysWallet{}
	wallet.Test = false
	privateKey := "AhSeFuxHydXkDV3Qp1F1onoPTsG6fFeLMsB9K7qBtjwM"
	wallet.FromPrivateKey(privateKey)
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//AhSeFuxHydXkDV3Qp1F1onoPTsG6fFeLMsB9K7qBtjwM
	//CXwmA9bYf6AP53jp2jzRG5nCtV8QEPe6AwbbX9a2dM6n
	//AR7EFpUF2bmztKLNPbEL8CT5HpahyiWwwo8
}

func TestVsysWallet_FromPublicKey(t *testing.T) {
	wallet := VsysWallet{}
	wallet.Test = false
	publicKey := "CXwmA9bYf6AP53jp2jzRG5nCtV8QEPe6AwbbX9a2dM6n"
	wallet.FromPublicKey(publicKey)
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//CXwmA9bYf6AP53jp2jzRG5nCtV8QEPe6AwbbX9a2dM6n
	//AR7EFpUF2bmztKLNPbEL8CT5HpahyiWwwo8
}

func TestVsysWallet_FromMnemonic(t *testing.T) {
	wallet := VsysWallet{}
	wallet.Test = false
	mnemonic := "inherit fiction obscure angry scale mouse fine tornado fitness before cricket parade"
	wallet.FromMnemonic(mnemonic)
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//inherit fiction obscure angry scale mouse fine tornado fitness before cricket parade
	//AhSeFuxHydXkDV3Qp1F1onoPTsG6fFeLMsB9K7qBtjwM
	//CXwmA9bYf6AP53jp2jzRG5nCtV8QEPe6AwbbX9a2dM6n
	//AR7EFpUF2bmztKLNPbEL8CT5HpahyiWwwo8
}
