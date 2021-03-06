module github.com/andnux/go-wallet

go 1.14

require (
	github.com/FactomProject/basen v0.0.0-20150613233007-fe3947df716e // indirect
	github.com/FactomProject/btcutilecc v0.0.0-20130527213604-d3a63a5752ec // indirect
	github.com/FactomProject/go-bip32 v0.3.5 // indirect
	github.com/FactomProject/go-bip39 v0.3.5
	github.com/FactomProject/go-bip44 v0.0.0-20190306062959-b541a96d8da9
	github.com/binance-chain/go-sdk v1.2.3
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcutil v1.0.2
	github.com/cmars/basen v0.0.0-20150613233007-fe3947df716e // indirect
	github.com/cosmos/cosmos-sdk v0.39.0
	github.com/dvsekhvalnov/jose2go v0.0.0-20180829124132-7f401d37b68a
	github.com/eoscanada/eos-go v0.9.0
	github.com/ethereum/go-ethereum v1.9.17
	github.com/filecoin-project/go-crypto v0.0.0-20191218222705-effae4ea9f03
	github.com/gcash/bchd v0.16.5
	github.com/gcash/bchutil v0.0.0-20200229194731-128fc9884722
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20190812055157-5d271430af9f // indirect
	github.com/ipfs/go-cid v0.0.5 // indirect
	github.com/ipfs/go-ipld-cbor v0.0.4
	github.com/ipfs/go-ipld-format v0.0.2 // indirect
	github.com/ltcsuite/ltcd v0.20.1-beta
	github.com/ltcsuite/ltcutil v0.0.0-20191227053721-6bec450ea6ad
	github.com/magiconair/properties v1.8.1
	github.com/minio/blake2b-simd v0.0.0-20160723061019-3f5f724cb5b1
	github.com/minio/sha256-simd v0.1.1
	github.com/mr-tron/base58 v1.1.3
	github.com/multiformats/go-varint v0.0.5
	github.com/nervosnetwork/ckb-sdk-go v0.1.0
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/polydawn/refmt v0.0.0-20190809202753-05966cbd336a
	github.com/rubblelabs/ripple v0.0.0-20200627211644-1ecb0c494a6a
	github.com/sasaxie/go-client-api v0.0.0-20190820063117-f0587df4b72e
	github.com/schancel/cashaddr-converter v0.0.0-20181111022653-4769e7add95a
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/go-amino v0.15.1
	github.com/tendermint/tendermint v0.33.6
	github.com/walkbean/vsys-sdk-go v0.0.0-20200425021120-95c96ca86edd
	github.com/warpfork/go-wish v0.0.0-20190328234359-8b3e70f8e830 // indirect
	github.com/whyrusleeping/cbor-gen v0.0.0-20200414195334-429a0b5e922e
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	golang.org/x/mobile v0.0.0-20200721161523-bcce01171201 // indirect
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1 // indirect
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	launchpad.net/gocheck v0.0.0-20140225173054-000000000087 // indirect
)

replace github.com/tendermint/go-amino => github.com/binance-chain/bnc-go-amino v0.14.1-binance.1
