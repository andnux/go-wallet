package go_wallet

import (
	"encoding/hex"
	"fmt"
	"github.com/andnux/go-wallet/filcoin/address"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"testing"
)

func TestAtomWallet_BuildFromPublicKey(t *testing.T) {
	wallet := AtomWallet{}
	hexByte := hex.EncodeToString([]byte("你哈"))
	wallet.BuildFromPublicKey(hexByte)
	fmt.Printf(wallet.GetAddress())
}

func TestAtomWallet_BuildFromRandomGenerate(t *testing.T) {
	wallet := AtomWallet{}
	wallet.BuildFromRandomGenerate()
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetAddress())
}

func TestName(t *testing.T) {

}
