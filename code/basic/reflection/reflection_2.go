/**
 * 反射获取结构中的匿名字段
 */
package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	User
	title string
}

func main() {
	m := Manager{User: User{1, "ok", 12}, title: "123"}
	t := reflect.TypeOf(m)

	//通过索引获取匿名字段，输出：reflect.StructField{Name:"User", PkgPath:"", Type:(*reflect.rtype)(0x10ae840), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:true}
	fmt.Printf("%#v\n", t.Field(0))

	//获取匿名字段中的字段,输出reflect.StructField{Name:"Id", PkgPath:"", Type:(*reflect.rtype)(0x109ee00), Tag:"", Offset:0x0, Index:[]int{0}, Anonymous:false}
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
}
