package xrp

import (
	"fmt"
	"testing"
)

func TestXrpWallet_BuildFromPublicKey(t *testing.T) {
	wallet := XrpWallet{}
	wallet.Test = false
	key := "ED9434799226374926EDA3B54B1B461B4ABF7237962EAE18528FEA67595397FA32"
	wallet.BuildFromPublicKey(key)
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
}
