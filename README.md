#### 1、Android SDK 使用
```shell script
gomobile bind -target android github.com/andnux/go-wallet
or 
cd github.com/andnux/go-wallet
gomobile bind -target android . ./atom ./bch ./btc ./ckb ./eos ./eth ./filcoin ./ltc ./omni ./tron ./vsys
```
![android.png](https://upload-images.jianshu.io/upload_images/3190592-25fd615c53e1052d.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)
#### 2、IOS SDK 使用
```shell script
gomobile bind -target ios github.com/andnux/go-wallet
or 
cd github.com/andnux/go-wallet
gomobile bind -target ios . ./atom ./bch ./btc ./ckb ./eos ./eth ./filcoin ./ltc ./omni ./tron ./vsys
```
#### 3、GO SDK 使用
```shell script
gomobile bind -target ios github.com/andnux/go-wallet
```
#### 4、目前支持： BTC EOS ETH FileCoin OMNI TRON VSYS
1、随机生成账号
```go
wallet = BtcWallet{}// EosWallet EthWallet FileCoinWallet TronWallet OmniWallet VsysWallet
wallet.Test = false
wallet.FromGenerate()
fmt.Println(wallet.GetPrivateKey())
fmt.Println(wallet.GetPublicKey())
fmt.Println(wallet.GetMnemonic())
fmt.Println(wallet.GetAddress())

//5JDawoh8AdLvhUZ7P7yT2qRsMtGoma7kFWgNTcELy48mCEirCPt
//049481269601b31efedd81eb59293558d7684ebaf129987a724e3ba692cfd8750ac6ea525f25f0dded594fa29f6041536deaaddb56d34f97b4ec1fd5f08e4d3d45
//undo dynamic dust become chat cage pool junk sphere next rent creek
//1NFheqXgVf78vU1sMRXjdbiqyjtQMLF5Vz
```
2、私钥创建账号
```go
wallet := BtcWallet{}// EosWallet EthWallet FileCoinWallet TronWallet OmniWallet VsysWallet
wallet.Test = false
wallet.FromPrivateKey("5JDawoh8AdLvhUZ7P7yT2qRsMtGoma7kFWgNTcELy48mCEirCPt")
fmt.Println(wallet.GetPrivateKey())
fmt.Println(wallet.GetPublicKey())
fmt.Println(wallet.GetMnemonic())
fmt.Println(wallet.GetAddress())
//5JDawoh8AdLvhUZ7P7yT2qRsMtGoma7kFWgNTcELy48mCEirCPt
//049481269601b31efedd81eb59293558d7684ebaf129987a724e3ba692cfd8750ac6ea525f25f0dded594fa29f6041536deaaddb56d34f97b4ec1fd5f08e4d3d45
//
//1NFheqXgVf78vU1sMRXjdbiqyjtQMLF5Vz
```
3、助记词创建账号
```go
wallet := BtcWallet{}// EosWallet EthWallet FileCoinWallet TronWallet OmniWallet VsysWallet
wallet.Test = false
mnemonic := "undo dynamic dust become chat cage pool junk sphere next rent creek"
wallet.FromMnemonic(mnemonic)
fmt.Println(wallet.GetPrivateKey())
fmt.Println(wallet.GetPublicKey())	
fmt.Println(wallet.GetMnemonic())
fmt.Println(wallet.GetAddress())
//5JDawoh8AdLvhUZ7P7yT2qRsMtGoma7kFWgNTcELy48mCEirCPt
//049481269601b31efedd81eb59293558d7684ebaf129987a724e3ba692cfd8750ac6ea525f25f0dded594fa29f6041536deaaddb56d34f97b4ec1fd5f08e4d3d45	
//undo dynamic dust become chat cage pool junk sphere next rent creek
//1NFheqXgVf78vU1sMRXjdbiqyjtQMLF5Vz
```
4、助记词通过路径生成账号
```go
wallet := BtcWallet{}// EosWallet EthWallet FileCoinWallet TronWallet OmniWallet VsysWallet
wallet.Test = false
mnemonic := "undo dynamic dust become chat cage pool junk sphere next rent creek"
wallet.FromMnemonicAndPath(mnemonic,"m/44'/200'/0'/0/0")
fmt.Println(wallet.GetPrivateKey())
fmt.Println(wallet.GetPublicKey())
fmt.Println(wallet.GetMnemonic())
fmt.Println(wallet.GetAddress())
//5JDawoh8AdLvhUZ7P7yT2qRsMtGoma7kFWgNTcELy48mCEirCPt	
//049481269601b31efedd81eb59293558d7684ebaf129987a724e3ba692cfd8750ac6ea525f25f0dded594fa29f6041536deaaddb56d34f97b4ec1fd5f08e4d3d45
//undo dynamic dust become chat cage pool junk sphere next rent creek
//1NFheqXgVf78vU1sMRXjdbiqyjtQMLF5Vz
```
5、签名
```go
wallet := BtcWallet{}// EosWallet EthWallet FileCoinWallet TronWallet OmniWallet VsysWallet
wallet.Test = false
mnemonic := "undo dynamic dust become chat cage pool junk sphere next rent creek"
wallet.FromMnemonicAndPath(mnemonic,"m/44'/200'/0'/0/0")
signed:=wallet.Sign(data) //签名
```
