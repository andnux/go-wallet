package btc

import (
	"fmt"
	"testing"
)

func TestBtcWallet_BuildFromRandomGenerate(t *testing.T) {
	wallet := BtcWallet{}
	wallet.Test = false
	wallet.BuildFromRandomGenerate()
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	//5JDawoh8AdLvhUZ7P7yT2qRsMtGoma7kFWgNTcELy48mCEirCPt
	//049481269601b31efedd81eb59293558d7684ebaf129987a724e3ba692cfd8750ac6ea525f25f0dded594fa29f6041536deaaddb56d34f97b4ec1fd5f08e4d3d45
	//undo dynamic dust become chat cage pool junk sphere next rent creek
	//1NFheqXgVf78vU1sMRXjdbiqyjtQMLF5Vz
}

func TestBtcWallet_BuildFromPrivateKey(t *testing.T) {
	wallet := BtcWallet{}
	wallet.Test = false
	wallet.BuildFromPrivateKey("5JDawoh8AdLvhUZ7P7yT2qRsMtGoma7kFWgNTcELy48mCEirCPt")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	//5JDawoh8AdLvhUZ7P7yT2qRsMtGoma7kFWgNTcELy48mCEirCPt
	//049481269601b31efedd81eb59293558d7684ebaf129987a724e3ba692cfd8750ac6ea525f25f0dded594fa29f6041536deaaddb56d34f97b4ec1fd5f08e4d3d45
	//
	//1NFheqXgVf78vU1sMRXjdbiqyjtQMLF5Vz
}

func TestBtcWallet_BuildFromMnemonic(t *testing.T) {
	wallet := BtcWallet{}
	wallet.Test = false
	mnemonic := "undo dynamic dust become chat cage pool junk sphere next rent creek"
	wallet.BuildFromMnemonic(mnemonic)
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	//5JDawoh8AdLvhUZ7P7yT2qRsMtGoma7kFWgNTcELy48mCEirCPt
	//049481269601b31efedd81eb59293558d7684ebaf129987a724e3ba692cfd8750ac6ea525f25f0dded594fa29f6041536deaaddb56d34f97b4ec1fd5f08e4d3d45
	//undo dynamic dust become chat cage pool junk sphere next rent creek
	//1NFheqXgVf78vU1sMRXjdbiqyjtQMLF5Vz
}

func TestBtcWallet_BuildFromMnemonicAndPath(t *testing.T) {
	wallet := BtcWallet{}
	wallet.Test = false
	mnemonic := "undo dynamic dust become chat cage pool junk sphere next rent creek"
	wallet.BuildFromMnemonicAndPath(mnemonic, "m/44'/200'/0'/0/0")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	//5JDawoh8AdLvhUZ7P7yT2qRsMtGoma7kFWgNTcELy48mCEirCPt
	//049481269601b31efedd81eb59293558d7684ebaf129987a724e3ba692cfd8750ac6ea525f25f0dded594fa29f6041536deaaddb56d34f97b4ec1fd5f08e4d3d45
	//undo dynamic dust become chat cage pool junk sphere next rent creek
	//1NFheqXgVf78vU1sMRXjdbiqyjtQMLF5Vz
}

func TestBtcWallet_BuildFromPublicKey(t *testing.T) {
	wallet := BtcWallet{}
	wallet.Test = false
	wallet.BuildFromPublicKey("049481269601b31efedd81eb59293558d7684ebaf129987a724e3ba692cfd8750ac6ea525f25f0dded594fa29f6041536deaaddb56d34f97b4ec1fd5f08e4d3d45")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	//
	//049481269601b31efedd81eb59293558d7684ebaf129987a724e3ba692cfd8750ac6ea525f25f0dded594fa29f6041536deaaddb56d34f97b4ec1fd5f08e4d3d45
	//
	//1NFheqXgVf78vU1sMRXjdbiqyjtQMLF5Vz
}
