package go_wallet

import (
	"fmt"
	"testing"
)

func TestOmniWallet_BuildFromRandomGenerate(t *testing.T) {
	wallet := OmniWallet{}
	wallet.Test = false
	wallet.BuildFromRandomGenerate()
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	//5JQZimWpQnL77kLnH5D57mzBtsjbN8KEJUAznmWpUZdaVnNTz52
	//0467b178cb985e77044bc33336116c6d79130b94e173fee4a7be1c1a8bb3c102ddd16436045d05260b1f3217269a688fa9c789fde27acbff3d8fba68a16199398e
	//convince actor eyebrow marine all faculty meat draft viable scan bridge misery
	//16UzjrHxMCgFFYzwUmGM98uTY3rgwU6wE6
}

func TestOmniWallet_BuildFromPrivateKey(t *testing.T) {
	wallet := OmniWallet{}
	wallet.Test = false
	wallet.BuildFromPrivateKey("5JQZimWpQnL77kLnH5D57mzBtsjbN8KEJUAznmWpUZdaVnNTz52")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	//5JQZimWpQnL77kLnH5D57mzBtsjbN8KEJUAznmWpUZdaVnNTz52
	//0467b178cb985e77044bc33336116c6d79130b94e173fee4a7be1c1a8bb3c102ddd16436045d05260b1f3217269a688fa9c789fde27acbff3d8fba68a16199398e
	//
	//16UzjrHxMCgFFYzwUmGM98uTY3rgwU6wE6
}

func TestOmniWallet_BuildFromMnemonic(t *testing.T) {
	wallet := OmniWallet{}
	wallet.Test = false
	mnemonic := "convince actor eyebrow marine all faculty meat draft viable scan bridge misery"
	wallet.BuildFromMnemonic(mnemonic)
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	//5JQZimWpQnL77kLnH5D57mzBtsjbN8KEJUAznmWpUZdaVnNTz52
	//0467b178cb985e77044bc33336116c6d79130b94e173fee4a7be1c1a8bb3c102ddd16436045d05260b1f3217269a688fa9c789fde27acbff3d8fba68a16199398e
	//convince actor eyebrow marine all faculty meat draft viable scan bridge misery
	//16UzjrHxMCgFFYzwUmGM98uTY3rgwU6wE6
}

func TestOmniWallet_BuildFromPublicKey(t *testing.T) {
	wallet := OmniWallet{}
	wallet.Test = false
	wallet.BuildFromPublicKey("0467b178cb985e77044bc33336116c6d79130b94e173fee4a7be1c1a8bb3c102ddd16436045d05260b1f3217269a688fa9c789fde27acbff3d8fba68a16199398e")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	//
	//0467b178cb985e77044bc33336116c6d79130b94e173fee4a7be1c1a8bb3c102ddd16436045d05260b1f3217269a688fa9c789fde27acbff3d8fba68a16199398e
	//
	//16UzjrHxMCgFFYzwUmGM98uTY3rgwU6wE6
}
