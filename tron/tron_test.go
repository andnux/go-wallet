package tron

import (
	"fmt"
	"testing"
)

func TestTronWallet_BuildFromRandomGenerate(t *testing.T) {
	wallet := TronWallet{}
	wallet.Test = false
	wallet.BuildFromRandomGenerate()
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//collect omit kid pond country result cabin core inquiry sand differ tuition
	//1d19584d58679aebe81f2eb2e47b1ae21d7615af874e20374dd91798de9eef31
	//027aff400c1dcd212de84fca92a983547603061f8beb55d7c1725fb2ebc82999b7
	//41c6f59bd0b8e11acd5ae6eba29c3966a656c22f8a
}

func TestTronCoinWallet_GetPrivateKey(t *testing.T) {
	wallet := TronWallet{}
	wallet.Test = false
	wallet.BuildFromPrivateKey("1d19584d58679aebe81f2eb2e47b1ae21d7615af874e20374dd91798de9eef31")
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//
	//1d19584d58679aebe81f2eb2e47b1ae21d7615af874e20374dd91798de9eef31
	//027aff400c1dcd212de84fca92a983547603061f8beb55d7c1725fb2ebc82999b7
	//41c6f59bd0b8e11acd5ae6eba29c3966a656c22f8a
}

func TestTronWallet_BuildFromMnemonic(t *testing.T) {
	wallet := TronWallet{}
	wallet.Test = false
	wallet.BuildFromMnemonic("collect omit kid pond country result cabin core inquiry sand differ tuition")
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//collect omit kid pond country result cabin core inquiry sand differ tuition
	//1d19584d58679aebe81f2eb2e47b1ae21d7615af874e20374dd91798de9eef31
	//027aff400c1dcd212de84fca92a983547603061f8beb55d7c1725fb2ebc82999b7
	//41c6f59bd0b8e11acd5ae6eba29c3966a656c22f8a
}

func TestTronWallet_BuildFromPublicKey(t *testing.T) {
	wallet := TronWallet{}
	wallet.Test = false
	wallet.BuildFromPublicKey("027aff400c1dcd212de84fca92a983547603061f8beb55d7c1725fb2ebc82999b7")
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	fmt.Println(wallet.ConvertToTAddress())
	//
	//
	//027aff400c1dcd212de84fca92a983547603061f8beb55d7c1725fb2ebc82999b7
	//41c6f59bd0b8e11acd5ae6eba29c3966a656c22f8a
	//TU7D2EHTdSW4cTo71rrhE8Bq48LjXdymyk
}

func TestTronWallet_BuildFromAddress(t *testing.T) {
	wallet := TronWallet{}
	wallet.Test = false
	wallet.BuildFromAddress("TU7D2EHTdSW4cTo71rrhE8Bq48LjXdymyk")
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	fmt.Println(wallet.ConvertTo4Address())
	//
	//
	//TU7D2EHTdSW4cTo71rrhE8Bq48LjXdymyk
	//41c6f59bd0b8e11acd5ae6eba29c3966a656c22f8a
}
