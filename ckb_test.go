package go_wallet

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/nervosnetwork/ckb-sdk-go/address"
	"github.com/nervosnetwork/ckb-sdk-go/transaction"
	"github.com/nervosnetwork/ckb-sdk-go/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCkbWallet_BuildFromRandomGenerate(t *testing.T) {
	wallet := CkbWallet{}
	wallet.Test = false
	wallet.BuildFromRandomGenerate()
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
}

func TestCkbWallet_BuildFromPrivateKey(t *testing.T) {
	wallet := CkbWallet{}
	wallet.Test = false
	wallet.BuildFromPrivateKey("40d9cc3facc0aaeb6e4c2783d80b9fd72ac19f15f218946ab73fd8a19b54d3e6")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "020fefb65378b406c6c445ef8b34030fa75f8da25b83f1de5aa543835d9dfbf1fd", wallet.GetPublicKey())
	assert.Equal(t, "ckb1qyqgmajgkkwsy6nqz4fu8axg6fg88znmjgqqxusdpg", wallet.GetAddress())
}

func TestCkbWallet_BuildFromMnemonic(t *testing.T) {
	wallet := CkbWallet{}
	wallet.Test = false
	mnemonic := "paper disagree someone pioneer ball latin voyage remain add slot double loan"
	wallet.BuildFromMnemonic(mnemonic)
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "40d9cc3facc0aaeb6e4c2783d80b9fd72ac19f15f218946ab73fd8a19b54d3e6", wallet.GetPrivateKey())
}

func TestName(t *testing.T) {
	script := &types.Script{
		CodeHash: types.HexToHash(transaction.SECP256K1_BLAKE160_SIGHASH_ALL_TYPE_HASH),
		HashType: types.HashTypeType,
		Args:     common.Hex2Bytes("b39bbc0b3673c7d36450bc14cfcdad2d559c6c64"),
	}
	addr, err := address.Generate(address.Mainnet, script)
	if err != nil {
		panic(err)
		return
	}
	assert.Equal(t, addr, "ckb1qyqt8xaupvm8837nv3gtc9x0ekkj64vud3jqfwyw5v")
}

func TestCkbWallet_BuildFromMnemonicAndPath(t *testing.T) {
	wallet := CkbWallet{}
	wallet.Test = false
	mnemonic := "paper disagree someone pioneer ball latin voyage remain add slot double loan"
	wallet.BuildFromMnemonicAndPath(mnemonic, "m/44'/309'/0'/0/0")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "40d9cc3facc0aaeb6e4c2783d80b9fd72ac19f15f218946ab73fd8a19b54d3e6", wallet.GetPrivateKey())
	assert.Equal(t, "020fefb65378b406c6c445ef8b34030fa75f8da25b83f1de5aa543835d9dfbf1fd", wallet.GetPublicKey())
	//assert.Equal(t,"ckb1qyqgmajgkkwsy6nqz4fu8axg6fg88znmjgqqxusdpg",wallet.GetAddress())
}

func TestCkbWallet_BuildFromPublicKey(t *testing.T) {
	wallet := CkbWallet{}
	wallet.Test = false
	wallet.BuildFromPublicKey("020fefb65378b406c6c445ef8b34030fa75f8da25b83f1de5aa543835d9dfbf1fd")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "ckb1qyqgmajgkkwsy6nqz4fu8axg6fg88znmjgqqxusdpg", wallet.GetAddress())
}
