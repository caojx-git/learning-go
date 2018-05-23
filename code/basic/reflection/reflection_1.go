/**
 * 反射获取接结构中的属性和方法
 *
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

func (u User) Hello() {
	fmt.Println("hello world")
}

func main() {
	u := User{1, "ok", 12}
	Info(u)
}

func Info(o interface{}) {
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	v := reflect.ValueOf(o)
	fmt.Println("Fields:")

	//1.通过反射输出所有的属性和值
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s: %v = %v\n", f.Name, f.Type, val)
	}

	//2.通过反射获取所有的方法
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}
}
