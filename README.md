# prolicense

私有Key：`Lsdiuqwjasksdiasdlwoerfasdwerfas`

## 加密成key

```go
//CheckType  string `json:"check_type"` //验证的方式，1.程序运行验证的方式  2.返回个数的验证方式
//CheckMac   string `json:"check_mac"`
//LicenseKey string `json:"license_key"` //key
//EndTime    int8   `json:"end_time"`    //key的过期时间戳
//SipNumber  int    `json:"sip_number"`  //sip的数量

var li = pojo.LicenseModel{}
li.SipNumber = 1  //限制的许可数量
li.EndTime = 2

--传入私有的key
li.GetKey("Lsdiuqwjasksdiasdlwoerfasdwerfas")
返回key：
li.LicenseKey
```

## 验证key

```go
var li = pojo.LicenseModel{}
f :=li.CheckKey(publicKey,privateKey) //查找限制的数量
jsonErr:=json.Unmarshal([]byte(f),&li)
```





# 处理

http://localhost:6600/MAC  获取mac的key的服务端

```go
通过prolicense调用

//加密返回key给客户端
	var li = pojo.LicenseModel{}
	li.SipNumber = 1  //填写数量
	li.EndTime = 2
	//如果mac地址为空的话，就是本机的mac地址
	li.CheckMac = "00:ff:31:a9:1c:a000:ff:28:c0:3e:1a28:d0:ea:5c:9b:492a:d0:ea:5c:9b:4938:f3:ab:43:c1:3600:50:56:c0:00:0100:50:56:c0:00:0828:d0:ea:5c:9b:4d00:15:5d:4b:09:06"
	li.GetKey("Lsdiuqwjasksdiasdlwoerfasdwerfas") //私有的key值

	fmt.Printf("本机MAC地址：%s\nkey：%s\n过期时间:%d\n数量：%d\n", li.CheckMac, li.LicenseKey, li.EndTime, li.SipNumber)
```

返回`li.LicenseKey`，加密后的key放入到验证项目的配置文件中。

然后通过`私有key`和`加密key` 来解析出对应的数据

