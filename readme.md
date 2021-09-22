# goclub/type

```shell
go get github.com/goclub/type
```

go 的类型系统没有 int? 这样的语法来表达一些基础字段是 nil 还是zero value.有些业务场景需要实现类似 int? 来解决问题.

goclub/type 提供了 Option 类型来解决这类问题. 

```go
package main

import (
	xtype "github.com/goclub/type"
	"log"
)

// 利用 xtype.OptionInt 判断是否设置 age
func SetData(name string, age xtype.OptionInt) {
	if age.Valid() {
		log.Print("set age", age.Unwrap())
	} else {
		log.Print("not set age")
	}
}

func main() {
	SetData("goclub", xtype.Int(18)) // set age 18
	SetData("goclub", xtype.OptionInt{}) // not set age
}
```

xtype.OptionInt源码:

```go
package xtype
type OptionInt struct {
	int int
	valid bool
}
func (o OptionInt) Valid() bool {
	return o.valid
}
func (o OptionInt) Unwrap() int {
	if o.valid {
		return o.int
	}
	return 0
}
func Int(i int) OptionInt {
	return OptionInt{ i, true}
}
```

其他基础类型也都实现了 Option 方法,通过 `xtype.类型()` 获取有效值, `xtype.Option类型{}` 表示无效值. 