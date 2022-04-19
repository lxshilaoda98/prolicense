package main

import (
	"fmt"
	"github.com/lxshilaoda98/prolicense/pojo"
)

//s struct {
//Code int    `json:"code"` //返回代码 0：成功 1：失败
//Msg  string `json:"msg"`  //返回描述
//}
//CheckType  string `json:"check_type"` //验证的方式，1.程序运行验证的方式  2.返回个数的验证方式
//CheckMac   string `json:"check_mac"`
//LicenseKey string `json:"license_key"` //key
//EndTime    int8   `json:"end_time"`    //key的过期时间戳
//SipNumber  int    `json:"sip_number"`  //sip的数量
func main() {
	var li = pojo.LicenseModel{}
	li.SipNumber = 500
	li.EndTime = 2
	li.GetKey("Lsdiuqwjasksdiasdlwoerfasdwerfas")
	fmt.Printf("本机MAC地址：%s\nkey：%s\n过期时间:%d\n数量：%d\n", li.CheckMac, li.LicenseKey, li.EndTime, li.SipNumber)

	//
	////li.CheckKey("asdiuqwjasksdiasdlwoerfasdwerfas")
	////传入li的时间。如果为0就是永不过期1
	//
	//li.EndTime = 0
	//li.GetKey("Lsdiuqwjasksdiasdlwoerfasdwerfas") //通过公有key获取private的key
	//fmt.Println(li.LicenseKey)
	//
	////通用key
	////验证key是否正确
	f := li.CheckKey(li.LicenseKey,
		"Lsdiuqwjasksdiasdlwoerfasdwerfas")

	fmt.Println("f:", f)

}
