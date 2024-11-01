// Description: Go 语言学习测试代码
// go mod tidy 生成 go.sum 文件
// go mod init <module-name> 生成 go.mod 文件
// go get <module-name> 下载依赖包
// go get -u <module-name> 更新依赖包

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"

	"github.com/shirou/gopsutil/mem"
)

var value int = 1

func add(a int, b int) int {
	return a + b
}
func operatingSystem() string {
	hostname, err := os.Hostname()
	// nil 是一个预先声明的标识符，用于表示零值、空值或者是一个空指针 无错误
	if err == nil {
		fmt.Println("主机名：", hostname)
	}
	gopath := os.Getenv("GOPATH")
	fmt.Println("go环境变量GOPATH: ", gopath)
	fmt.Println("系统架构: ", runtime.GOARCH)
	// 声明一个 runtime.MemStats 类型的变量 memstats
	// runtime.memsats 是一个结构体类型，用于存储内存分配的统计信息
	var memstats runtime.MemStats
	runtime.ReadMemStats(&memstats)
	fmt.Print("Alloc: ", memstats.Alloc/1024/1024, "MB")
	fmt.Print("\t TotalAlloc: ", memstats.TotalAlloc/1024/1024, "MB")
	fmt.Print("\t Sys: ", memstats.Sys/1024/1024, "MB")
	fmt.Printf("\t NumGC: %v \n", memstats.NumGC)
	// 获取系统内存
	v, err := mem.VirtualMemory()
	if err == nil {
		fmt.Println("系统内存：", v.Total/1024/1024/1024, "GB")
	}

	return runtime.GOOS
}

// 测试函数
func test1() {
	// 变量和赋值
	a0, b0 := 1, 2
	value = add(a0, b0)
	fmt.Println("a + b", "= ", value)
	bitvalue := 1
	fmt.Println("位运算左移 ", bitvalue<<1)
	fmt.Println("位运算右移 ", bitvalue>>1)

	// iota 常量生成器
	const (
		a = iota
		b
		c
		d = 1 << iota
		e
	)
	fmt.Println(a, b, c, d, e)
	// if 条件语句
	if f := d * e; f <= 100 {
		fmt.Println("f < 10 ")
	} else {
		fmt.Println("f = ", f)
	}
	// for 循环语句
	imax := 10
	for i := 0; i < imax; i++ {
		if i == 0 {
			fmt.Printf("i = %d ", i)
			continue
		}
		fmt.Printf(" %d ", i)
		if i == imax-1 {
			fmt.Println()
		}
	}
	// switch 语句
	switch os := operatingSystem(); os {
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Other OS")
	}
}

func test2_btcutil() {
	// 生成私钥
	// SECP256K1
	privKey1, err := btcec.NewPrivateKey()
	if err != nil {
		fmt.Println("生成私钥失败")
		return
	}
	// 生成WIF格式的私钥
	wif1, err := btcutil.NewWIF(privKey1, &chaincfg.MainNetParams, true)
	if err != nil {
		fmt.Println("生成WIF失败")
		return
	}
	// 生成公钥 k * G = K
	pubKey1 := privKey1.PubKey()
	// P2PK Pay-to-PubKey
	// 用公钥接收和锁定资金 验证交易 早期使用

	// 生成P2PKH地址 Pay-to-PubKey-Hash 以 1 开头
	// 对公钥进行两次哈希 第一次哈希使用 SHA256 第二次哈希使用 RIPEMD160，生成公钥哈希 Public Key Hash
	// 在公钥哈希前加上版本号 0x00，生成版本化公钥哈希 Versioned Public Key Hash
	// 对版本化公钥哈希进行两次哈希 第一次哈希使用 SHA256 第二次哈希使用 SHA256，生成校验码 Checksum
	// 取校验码的前四个字节，加到版本化公钥哈希后面，生成校验版本化公钥哈希 Checksummed Versioned Public Key Hash
	// 对校验版本化公钥哈希进行 Base58 编码，生成 P2PKH 地址
	pubKeyHash1 := btcutil.Hash160(pubKey1.SerializeCompressed())
	addressPKH, err := btcutil.NewAddressPubKeyHash(pubKeyHash1, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println("生成P2PKH地址失败")
		return
	}

	// 生成P2SH地址 Pay-to-Script-Hash 以 3 开头
	// 常见用途是多重签名钱包 需要多个私钥签名才能完成交易的地址
	// 创建一个比特币脚本 常用OP_CHECKMULTISIG
	// 对脚本进行RIPEMD160哈希生成脚本哈希 Script Hash
	// 在脚本哈希前加上版本号 0x05，生成版本化脚本哈希 Versioned Script Hash
	// 对版本化脚本哈希进行两次哈希 第一次哈希使用 SHA256 第二次哈希使用 SHA256，生成校验码 Checksum
	// 取校验码的前四个字节，加到版本化脚本哈希后面，生成校验版本化脚本哈希 Checksummed Versioned Script Hash
	// 对校验版本化脚本哈希进行 Base58 编码，生成 P2SH 地址
	privKey2, err := btcec.NewPrivateKey()
	if err != nil {
		fmt.Println("生成私钥失败")
		return
	}
	// 生成多重签名脚本
	pubKey2 := []*btcec.PublicKey{privKey1.PubKey(), privKey2.PubKey()}
	var addressPubKeys []*btcutil.AddressPubKey
	for _, pubKey := range pubKey2 {
		addressPubKey, err := btcutil.NewAddressPubKey(pubKey.SerializeCompressed(), &chaincfg.MainNetParams)
		if err != nil {
			fmt.Println("生成地址公钥失败")
			return
		}
		addressPubKeys = append(addressPubKeys, addressPubKey)
	}
	script, err := txscript.MultiSigScript(addressPubKeys, 2)
	if err != nil {
		fmt.Println("生成多重签名脚本失败")
		return
	}
	// 生成P2SH地址
	addressP2SH, err := btcutil.NewAddressScriptHash(script, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println("生成P2SH地址失败")
		return
	}
	// 生成Bech32地址 bc1 开头 P2WPKH地址 Pay-to-Witness-PubKey-Hash
	// Hash160 里面做了两次哈希 一次是 SHA256 一次是 RIPEMD160
	pubKeyHash := btcutil.Hash160(pubKey1.SerializeCompressed())
	addressBech32, err := btcutil.NewAddressWitnessPubKeyHash(pubKeyHash, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Println("生成Bech32地址失败")
	}

	fmt.Println("私钥: ", privKey1.Serialize())
	fmt.Printf("私钥wif: %s \n", wif1.String())
	fmt.Printf("未压缩公钥： %x \n", pubKey1.SerializeUncompressed())
	fmt.Printf("压缩公钥： %x \n", pubKey1.SerializeCompressed())
	fmt.Println("P2PKH地址: ", addressPKH.EncodeAddress())
	fmt.Println("P2SH地址: ", addressP2SH.EncodeAddress())
	fmt.Println("Bech32地址: ", addressBech32.EncodeAddress())

}

func main() {
	test1()

	test2_btcutil()

}
