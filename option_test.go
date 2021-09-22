package xtype_test

import (
	xtype "github.com/goclub/type"
	"log"
	"testing"
)

// 必须设置 age
func SetDataV1(name string, age int) {
	log.Print("set age", age)
}
// 可通过 notSetAge 控制不设置 age
func SetDataV2(name string, age int, notSetAge bool) {
	if notSetAge {
		log.Print("not set age")
	} else {
		log.Print("set age", age)
	}
}
// 传递nil则表示不设置 age
func SetDataV3(name string, age *int) {
	if age == nil {
		log.Print("not set age")
	} else {
		log.Print("set age", *age) // // *age 才能取到 int 否则取到的是 *int
	}
}
// 利用 xtype.OptionInt 判断是否设置 age
func SetDataV4(name string, age xtype.OptionInt) {
	if age.Valid() {
		log.Print("set age", age.Unwrap())
	} else {
		log.Print("not set age")
	}
}

func ExampleInt() {
	SetDataV1("goclub", 0)  // set age 0
	SetDataV1("goclub", 18) // set age 18

	SetDataV2("goclub", 18, false) // set age 18
	SetDataV2("goclub", 0, true)   // not set age

	agev3 := 18
	SetDataV3("goclub", &agev3) // set age 18
	SetDataV3("goclub", nil) // not set age

	SetDataV4("goclub", xtype.Int(18)) // set age 18
	SetDataV4("goclub", xtype.OptionInt{}) // not set age
}
func TestOptionInt(t *testing.T) {
	ExampleInt()
}