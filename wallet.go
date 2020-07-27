package go_wallet

import (
	"bytes"
	"strconv"
	"strings"
)

type Bip44 struct {
	M            string
	Purpose      string
	CoinType     uint32
	Account      uint32
	Change       uint32
	AddressIndex uint32
}

func stringToUint32(str string) uint32 {
	num, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(num)
}

func bip44Parser(path string) (bip44 Bip44, err error) {
	s := strings.Split(path, "/")
	if len(s) != 6 {
		panic("path error")
		return bip44, nil
	}
	bip44 = Bip44{
		M:            strings.ReplaceAll(s[0], "'", ""),
		Purpose:      strings.ReplaceAll(s[1], "'", ""),
		CoinType:     stringToUint32(strings.ReplaceAll(s[2], "'", "")),
		Account:      stringToUint32(strings.ReplaceAll(s[3], "'", "")),
		Change:       stringToUint32(strings.ReplaceAll(s[4], "'", "")),
		AddressIndex: stringToUint32(strings.ReplaceAll(s[5], "'", "")),
	}
	return bip44, nil
}

func bytesCombine(pBytes ...[]byte) []byte {
	length := len(pBytes)
	s := make([][]byte, length)
	for index := 0; index < length; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

type Signature interface {
	Sign(data []byte) (signed []byte, err error)
}

type RandomGenerate interface {
	BuildFromRandomGenerate()
}

type Keystore interface {
	BuildFromKeystore(keystore string, password string)
	GetKeystore() string
}

type PrivateKey interface {
	BuildFromPrivateKey(privateKey string)
	GetPrivateKey() string
}

type PublicKey interface {
	BuildFromPublicKey(publicKey string)
	GetPublicKey() string
}

type Address interface {
	BuildFromAddress(address string)
	GetAddress() string
}

type Mnemonic interface {
	BuildFromMnemonicAndPath(mnemonic string, path string)
	BuildFromMnemonic(mnemonic string)
	GetMnemonic() string
}
