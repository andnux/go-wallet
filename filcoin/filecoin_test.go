package filcoin

import (
	"fmt"
	"testing"
)

func TestFileCoinWallet_FromGenerate(t *testing.T) {
	wallet := FileCoinWallet{}
	wallet.Test = false
	wallet.FromGenerate()
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//claw drill table grant off swamp spread amazing month aware write announce
	//c7d32f2f0d1b770ff183e56562cd839e9bceeb119ef8e7485e3b0e0de17078f7
	//04f83939f79e3f7be7c8e1e7032eb934b4902408bd48b9f6b0ccecfd0866a218cd957ffcc81aff483276874a96354200c31465fe3d595b8ca4f463f98c3b955c7c
	//f1w7bphex4rep5g6uw6cc5etq2jbfyusx7dqxfvny
}

func TestFileCoinWallet_FromPrivateKey(t *testing.T) {
	wallet := FileCoinWallet{}
	wallet.Test = false
	wallet.FromPrivateKey("c7d32f2f0d1b770ff183e56562cd839e9bceeb119ef8e7485e3b0e0de17078f7")
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//
	//c7d32f2f0d1b770ff183e56562cd839e9bceeb119ef8e7485e3b0e0de17078f7
	//04f83939f79e3f7be7c8e1e7032eb934b4902408bd48b9f6b0ccecfd0866a218cd957ffcc81aff483276874a96354200c31465fe3d595b8ca4f463f98c3b955c7c
	//f1w7bphex4rep5g6uw6cc5etq2jbfyusx7dqxfvny
}

func TestFileCoinWallet_FromMnemonic(t *testing.T) {
	wallet := FileCoinWallet{}
	wallet.Test = false
	wallet.FromMnemonic("claw drill table grant off swamp spread amazing month aware write announce")
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//claw drill table grant off swamp spread amazing month aware write announce
	//c7d32f2f0d1b770ff183e56562cd839e9bceeb119ef8e7485e3b0e0de17078f7
	//04f83939f79e3f7be7c8e1e7032eb934b4902408bd48b9f6b0ccecfd0866a218cd957ffcc81aff483276874a96354200c31465fe3d595b8ca4f463f98c3b955c7c
	//f1w7bphex4rep5g6uw6cc5etq2jbfyusx7dqxfvny
}

func TestFileCoinWallet_GetPublicKey(t *testing.T) {
	wallet := FileCoinWallet{}
	wallet.Test = false
	wallet.FromPublicKey("04f83939f79e3f7be7c8e1e7032eb934b4902408bd48b9f6b0ccecfd0866a218cd957ffcc81aff483276874a96354200c31465fe3d595b8ca4f463f98c3b955c7c")
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//
	//
	//04f83939f79e3f7be7c8e1e7032eb934b4902408bd48b9f6b0ccecfd0866a218cd957ffcc81aff483276874a96354200c31465fe3d595b8ca4f463f98c3b955c7c
	//f1w7bphex4rep5g6uw6cc5etq2jbfyusx7dqxfvny
}
