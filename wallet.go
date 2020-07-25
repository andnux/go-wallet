package go_wallet

type Account struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Address    string `json:"address"`
	Mnemonic   string `json:"mnemonic"`
	Keystore   string `json:"keystore"`
}

type Broadcast interface {
	Broadcast(data []byte) string
}

type Signature interface {
	Signature(data []byte, privateKey string) []byte
}

type Wallet interface {
	Name() string
	Generate(test bool) string
	GenerateByPrivateKey(privateKey string, test bool) string
	GenerateByMnemonic(mnemonic string, path string, test bool) string
}
