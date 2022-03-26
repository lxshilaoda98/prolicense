package main

import (
	"fmt"
	"github.com/n1n1n1_owner/prolicense/pojo"
)

func main() {
	var li = pojo.LicenseModel{}

	//li.CheckKey("asdiuqwjasksdiasdlwoerfasdwerfas")
	//传入li的时间。如果为0就是永不过期

	li.EndTime = 0
	li.GetKey("Lsdiuqwjasksdiasdlwoerfasdwerfas") //通过公有key获取private的key
	fmt.Println(li.LicenseKey)

	//通用key
	//验证key是否正确
	f := li.CheckKey(li.LicenseKey,
		"Lsdiuqwjasksdiasdlwoerfasdwerfas")
	fmt.Println("endtime:", li.EndTime)
	fmt.Println("f:", f)

}
