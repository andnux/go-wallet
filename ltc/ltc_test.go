package ltc

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestLtcWallet_BuildFromRandomGenerate(t *testing.T) {
	wallet := LtcWallet{}
	wallet.Test = false
	wallet.BuildFromRandomGenerate()
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	//6uLAGuze4uWXsojLdooiL7ogSWhVGm7ohxEb783ceF3Sqc4XUuF
	//0419db711a40e6985839de8efa2414bc77e47cba5a073aa76bc172e9ea23f966fdb1bc131f046b043fa2fc3e7289cf85f483170c2cdc7513436023d0e9186e0903
	//stereo clown donkey escape early exercise icon session exile modify glow danger
	//LMpz5A6sWcagA5imngDT173JSJSBTVWvZi
}

func TestLtcWallet_BuildFromPrivateKey(t *testing.T) {
	wallet := LtcWallet{}
	wallet.Test = false
	wallet.BuildFromPrivateKey("6uLAGuze4uWXsojLdooiL7ogSWhVGm7ohxEb783ceF3Sqc4XUuF")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "6uLAGuze4uWXsojLdooiL7ogSWhVGm7ohxEb783ceF3Sqc4XUuF", wallet.GetPrivateKey())
	assert.Equal(t, "0419db711a40e6985839de8efa2414bc77e47cba5a073aa76bc172e9ea23f966fdb1bc131f046b043fa2fc3e7289cf85f483170c2cdc7513436023d0e9186e0903", wallet.GetPublicKey())
	assert.Equal(t, "LMpz5A6sWcagA5imngDT173JSJSBTVWvZi", wallet.GetAddress())
}

func TestLtcWallet_BuildFromMnemonic(t *testing.T) {
	wallet := LtcWallet{}
	wallet.Test = false
	mnemonic := "stereo clown donkey escape early exercise icon session exile modify glow danger"
	wallet.BuildFromMnemonic(mnemonic)
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "6uLAGuze4uWXsojLdooiL7ogSWhVGm7ohxEb783ceF3Sqc4XUuF", wallet.GetPrivateKey())
	assert.Equal(t, "0419db711a40e6985839de8efa2414bc77e47cba5a073aa76bc172e9ea23f966fdb1bc131f046b043fa2fc3e7289cf85f483170c2cdc7513436023d0e9186e0903", wallet.GetPublicKey())
	assert.Equal(t, "LMpz5A6sWcagA5imngDT173JSJSBTVWvZi", wallet.GetAddress())
}

func TestLtcWallet_BuildFromMnemonicAndPath(t *testing.T) {
	wallet := LtcWallet{}
	wallet.Test = false
	mnemonic := "stereo clown donkey escape early exercise icon session exile modify glow danger"
	wallet.BuildFromMnemonicAndPath(mnemonic, "m/44'/2'/0'/0/0")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "6uLAGuze4uWXsojLdooiL7ogSWhVGm7ohxEb783ceF3Sqc4XUuF", wallet.GetPrivateKey())
	assert.Equal(t, "0419db711a40e6985839de8efa2414bc77e47cba5a073aa76bc172e9ea23f966fdb1bc131f046b043fa2fc3e7289cf85f483170c2cdc7513436023d0e9186e0903", wallet.GetPublicKey())
	assert.Equal(t, "LMpz5A6sWcagA5imngDT173JSJSBTVWvZi", wallet.GetAddress())
}

func TestLtcWallet_BuildFromPublicKey(t *testing.T) {
	wallet := LtcWallet{}
	wallet.Test = false
	wallet.BuildFromPublicKey("0419db711a40e6985839de8efa2414bc77e47cba5a073aa76bc172e9ea23f966fdb1bc131f046b043fa2fc3e7289cf85f483170c2cdc7513436023d0e9186e0903")
	fmt.Println(wallet.GetPrivateKey())
	fmt.Println(wallet.GetPublicKey())
	fmt.Println(wallet.GetMnemonic())
	fmt.Println(wallet.GetAddress())
	assert.Equal(t, "LMpz5A6sWcagA5imngDT173JSJSBTVWvZi", wallet.GetAddress())
}
