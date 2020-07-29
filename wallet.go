package go_wallet

import (
	"bytes"
	"github.com/tendermint/go-amino"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"strconv"
	"strings"
)

var MainCodec = amino.NewCodec()

const (
	privKeyAminoName = "tendermint/PrivKeySecp256k1"
	pubKeyAminoName  = "tendermint/PubKeySecp256k1"
)

func init() {
	MainCodec.RegisterInterface((*crypto.PubKey)(nil), nil)
	MainCodec.RegisterConcrete(secp256k1.PrivKeySecp256k1{},
		privKeyAminoName, nil)
	MainCodec.RegisterConcrete(secp256k1.PubKeySecp256k1{},
		pubKeyAminoName, nil)
}

type Bip44 struct {
	M            string
	Purpose      string
	CoinType     uint32
	Account      uint32
	Change       uint32
	AddressIndex uint32
}

func StringToUint32(str string) uint32 {
	num, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(num)
}

func Bip44Parser(path string) (bip44 Bip44, err error) {
	s := strings.Split(path, "/")
	if len(s) != 6 {
		panic("path error")
		return bip44, nil
	}
	bip44 = Bip44{}
	bip44.M = s[0]
	bip44.Purpose = strings.ReplaceAll(s[1], "'", "")
	if strings.Contains(s[2], "'") {
		bip44.CoinType = StringToUint32(strings.ReplaceAll(s[2], "'", "")) | 0x80000000
	} else {
		bip44.CoinType = StringToUint32(strings.ReplaceAll(s[2], "'", ""))
	}
	if strings.Contains(s[3], "'") {
		bip44.Account = StringToUint32(strings.ReplaceAll(s[3], "'", "")) | 0x80000000
	} else {
		bip44.Account = StringToUint32(strings.ReplaceAll(s[3], "'", ""))
	}
	if strings.Contains(s[4], "'") {
		bip44.Change = StringToUint32(strings.ReplaceAll(s[4], "'", "")) | 0x80000000
	} else {
		bip44.Change = StringToUint32(strings.ReplaceAll(s[4], "'", ""))
	}
	if strings.Contains(s[5], "'") {
		bip44.AddressIndex = StringToUint32(strings.ReplaceAll(s[5], "'", "")) | 0x80000000
	} else {
		bip44.AddressIndex = StringToUint32(strings.ReplaceAll(s[5], "'", ""))
	}
	return bip44, nil
}

func BytesCombine(pBytes ...[]byte) []byte {
	length := len(pBytes)
	s := make([][]byte, length)
	for index := 0; index < length; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

//下面代码IOS编译不过，暂时去掉
//type WalletSignature interface {
//	Sign(data []byte) (signed []byte)
//}

type Generate interface {
	FromGenerate()
}

type Keystore interface {
	FromKeystore(keystore string, password string)
	GetKeystore() string
}

type PrivateKey interface {
	FromPrivateKey(privateKey string)
	GetPrivateKey() string
}

type PublicKey interface {
	FromPublicKey(publicKey string)
	GetPublicKey() string
}

type Address interface {
	FromAddress(address string)
	GetAddress() string
}

type Mnemonic interface {
	FromMnemonicAndPath(mnemonic string, path string)
	FromMnemonic(mnemonic string)
	GetMnemonic() string
}
