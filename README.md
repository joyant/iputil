# iputil

`iputil`是一个`IP`帮助类库，提供`IPV4`和`IPV6`之间的转换、将字符串转成标准`IP`格式、简单的防火墙等功能，用`Go`语言编写。

## 工具函数

```
import (
    "github.com/joyant/iptuil"
)

iputil.IsIPV4("192.168.0.1") // true
iputil.IsIPV4("192.168.0.a") // false

iptuil.IsIPV6("::1") // true
iptuil.IsIPV6("1::") // true
iptuil.IsIPV6("1::1") // true
iptuil.IsIPV6("1008:124::1") // true
iptuil.IsIPV6("ABCD:EF01:2345:6789:ABCD:EF01:2345:6789") // true
iptuil.IsIPV6("1008:124::1::2") // false

iptuil.IsIP("::1") // true
iptuil.IsIP("192.168.0.1") // true
iptuil.IsIP("192.168.0.1000") // false

iputil.String2IPV4("192.168.0.1") // IPV4{192, 168, 0, 1}, true
iputil.String2IPV4("192.168.0.1000") // IPV4{0, 0, 0, 0}, false

iputil.String2IPV6("::1") // IPV6{0, 0, 0, 0, 0, 0, 0, 1}, true
iputil.String2IPV6(":::") // IPV6{0, 0, 0, 0, 0, 0, 0, 0}, false
iputil.String2IPV6("87:4B:2B:34::1") // IPV6{135, 75, 43, 52, 0, 0, 0, 1}, true

iputil.StringIPV42IPV6("192.168.0.1") // "::c0:a8:00:01", true
iputil.StringIPV42IPV6("192.168.0.999") // "", false
```

## 防火墙

```
f := NewFirewall()
err := f.LoadIP("192.168.0.*")
err := f.LoadIP("192.168.1-10.1")
err = f.LoadIP("a:b:c:e:f:1:2:3")
err = f.LoadIP("a:b:d:*::")

f.Match("192.168.0.10") // true
f.Match("192.168.9.1") // true
f.Match("a:b:c:e:f:1:2:3") // true
f.Match("a:b:d:ffff:0:0:0:0") // true
f.Match("12::") // false
```

## 许可证

Apache

## 参考

关于`IPV4`和`IPV6`概念，请(参考)["https://blog.csdn.net/Listen2You/article/details/98784656"]

