package go_wallet

type Account struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Address    string `json:"address"`
	Mnemonic   string `json:"mnemonic"`
	Keystore   string `json:"keystore"`
}

func InitEosWallet(test bool) *EosWallet {
	return &EosWallet{Test: test}
}

func InitVsysWallet(test bool) *VsysWallet {
	return &VsysWallet{Test: test}
}

func InitFileCoinWallet(test bool) *FileCoinWallet {
	return &FileCoinWallet{Test: test}
}

func InitEthWallet(test bool) *EthWallet {
	return &EthWallet{Test: test}
}

func InitBtcWallet(test bool) *BtcWallet {
	return &BtcWallet{Test: test}
}

//type Broadcast interface {
//	Broadcast(data []byte) string
//}
//
//type Signature interface {
//	Signature(data []byte, privateKey string) []byte
//}
//
//type Wallet interface {
//	Name() string
//	Generate() string
//	GenerateByPrivateKey(privateKey string) string
//	GenerateByMnemonic(mnemonic string, path string) string
//}
