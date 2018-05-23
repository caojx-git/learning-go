/**
 * 嵌入结构,嵌入结构成为组合，不是继承
 */
package main

import (
	"fmt"
)

type human struct {
	Sex int
}

type teacher struct {
	human
	Name string
	Age  int
}

type student struct {
	human
	Name string
	Age  int
}

func main() {
	a := teacher{Name: "joe", Age: 19, human: human{Sex: 0}}
	b := student{Name: "joe", Age: 20, human: human{Sex: 1}}

	a.Name = "joe"
	a.Age = 20

	//如果一个结构中嵌入可多个结构，且各个结构中存在同名字段，则可以使用下边的方式赋值
	a.human.Sex = 100
	//如果一个结构中嵌入可多个结构，且各个结构中不存在同名字段，则可以使用下边的方式赋值
	a.Sex = 101

	fmt.Println(a, b)
}
