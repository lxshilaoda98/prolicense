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

