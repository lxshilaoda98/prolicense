package pojo

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net"
)

type LicenseModel struct {
	s struct {
		Code int    `json:"code"` //返回代码 0：成功 1：失败
		Msg  string `json:"msg"`  //返回描述
	}
	CheckType  string `json:"check_type"` //验证的方式，1.程序运行验证的方式  2.返回个数的验证方式
	CheckMac   string `json:"check_mac"`
	LicenseKey string `json:"license_key"` //key
	EndTime    int8   `json:"end_time"`    //key的过期时间戳
	SipNumber  int    `json:"sip_number"`  //sip的数量
}

//验证key信息是否正确
func (l *LicenseModel) CheckKey(publicKey, key string) (retunKey string) {

	retunKey = AesDecrypt(publicKey, key) //给一个传过来的key和私有的一个key值

	return
}

//先获取加密的信息
//给一个过期的时间
func (l *LicenseModel) GetKey(key string) {
	//先获取本机的mac地址
	macStr := GetKey()
	fmt.Println(macStr)
	l.CheckMac = macStr
	l.s.Code = 0
	l.s.Msg = "获取成功"
	by, err := json.Marshal(l) //转成json格式
	resultKey := AesEncrypt(string(by), key)
	if err != nil {
		l.s.Code = 1
		l.s.Msg = "程序异常!传入的json不正确,请检查!"
	} else {
		l.LicenseKey = resultKey
	}
}

//region 公有方法
//加密
func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)

	// 分组秘钥
	block, err := aes.NewCipher(k)
	if err != nil {
		panic(fmt.Sprintf("key 长度必须 16/24/32长度: %s", err.Error()))
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	//使用RawURLEncoding 不要使用StdEncoding
	//不要使用StdEncoding  放在url参数中回导致错误
	return base64.RawURLEncoding.EncodeToString(cryted)

}

//通过mac地址获取key
func GetKey() (KeyStr string) {
	data := getMacAddrs()
	for _, v := range data {
		KeyStr += v
	}
	return
}
func getMacAddrs() (macAddrs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Printf("fail to get net interfaces: %v", err)
		return macAddrs
	}

	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}

		macAddrs = append(macAddrs, macAddr)
	}
	return macAddrs
}

//cryted   公有的一个key，约定key
//key 客户端传过来的key
func AesDecrypt(cryted string, key string) (ks string) {
	//使用RawURLEncoding 不要使用StdEncoding
	//不要使用StdEncoding  放在url参数中回导致错误
	crytedByte, err := base64.RawURLEncoding.DecodeString(cryted)

	k := []byte(key)

	// 分组秘钥
	block, err := aes.NewCipher(k)

	if err != nil {
		panic(fmt.Sprintf("key 长度必须 16/24/32长度: %s", err.Error()))
	}
	// 获取秘钥块的长度
	blockSize := block.BlockSize()

	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])

	// 创建数组
	orig := make([]byte, len(crytedByte))

	// 解密
	blockMode.CryptBlocks(orig, crytedByte)

	// 去补全码
	orig = PKCS7UnPadding(orig)

	ks = string(orig)
	return
}

//补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//endregion
