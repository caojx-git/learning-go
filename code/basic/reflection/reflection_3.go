/**
 * 通过反射修改字段内容
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

func main() {

	//基本类型
	x := 123
	v := reflect.ValueOf(&x)

	v.Elem().SetInt(999)

	fmt.Println(x)

	u := User{1, "ok", 12}
	Set(&u)
	fmt.Println(u)

}

func Set(o interface{}) {
	v := reflect.ValueOf(o)

	//是指针类型，并且是可以修改的
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("xxx")
		return
	} else {
		v = v.Elem()
	}

	f := v.FieldByName("Name")

	//没有找到对应的属性，打印BAD 返回
	if !f.IsValid() {
		fmt.Println("BAD")
		return
	}

	//找到Name属性，设置值为BYEBYE
	if f.Kind() == reflect.String {
		f.SetString("BYEBYE")
	}
}
