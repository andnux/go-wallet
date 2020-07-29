package eth

import (
	"fmt"
	"testing"
)

func TestEthWallet_BuildFromRandomGenerate(t *testing.T) {
	wallet := EthWallet{}
	wallet.Test = false
	wallet.BuildFromRandomGenerate()
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//inflict problem pipe senior security volcano cloth doll goose elephant provide expect
	//633b3d6ee8c02366a6fccbff8af9d0d4a0cd27130b8565baaa9311bd76a33ea4
	//0316c71e2943e5ed34be6ac2064a2396a1a28373bf6b694899d4a1ae0f2c951557
	//0xae550c1e3017daf94a66cb7668699c7ee450d9bf
}

func TestEthCoinWallet_GetPrivateKey(t *testing.T) {
	wallet := EthWallet{}
	wallet.Test = false
	wallet.BuildFromPrivateKey("633b3d6ee8c02366a6fccbff8af9d0d4a0cd27130b8565baaa9311bd76a33ea4")
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//
	//633b3d6ee8c02366a6fccbff8af9d0d4a0cd27130b8565baaa9311bd76a33ea4
	//0316c71e2943e5ed34be6ac2064a2396a1a28373bf6b694899d4a1ae0f2c951557
	//0xae550c1e3017daf94a66cb7668699c7ee450d9bf
}

func TestEthWallet_BuildFromMnemonic(t *testing.T) {
	wallet := EthWallet{}
	wallet.Test = false
	wallet.BuildFromMnemonic("inflict problem pipe senior security volcano cloth doll goose elephant provide expect")
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//inflict problem pipe senior security volcano cloth doll goose elephant provide expect
	//633b3d6ee8c02366a6fccbff8af9d0d4a0cd27130b8565baaa9311bd76a33ea4
	//0316c71e2943e5ed34be6ac2064a2396a1a28373bf6b694899d4a1ae0f2c951557
	//0xae550c1e3017daf94a66cb7668699c7ee450d9bf
}

func TestEthWallet_BuildFromPublicKey(t *testing.T) {
	wallet := EthWallet{}
	wallet.Test = false
	wallet.BuildFromPublicKey("0316c71e2943e5ed34be6ac2064a2396a1a28373bf6b694899d4a1ae0f2c951557")
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
	//
	//
	//0316c71e2943e5ed34be6ac2064a2396a1a28373bf6b694899d4a1ae0f2c951557
	//0xae550c1e3017daf94a66cb7668699c7ee450d9bf
}
