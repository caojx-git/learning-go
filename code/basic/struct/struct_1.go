/**
 * 定义struct结构
 */

package main

import (
	"fmt"
)

//定义一个结构
type test struct{}

type Person struct {
	Name string
	Age  int
}

func main() {
	a := test{}
	fmt.Println(a)

	//对结构中的字段赋值
	person := Person{}
	person.Age = 10
	person.Name = "joe"

	fmt.Println(person)

	person2 := Person{"tom", 19}
	fmt.Println(person2)

	person3 := Person{Age: 18, Name: "tom2"}
	fmt.Println(person3)

	//结构作为方法或函数的参数时，进行的是值拷贝
	A(person3)
	fmt.Println(person3)

	//可以通过指针解决上边值拷贝的问题
	B(&person3)
	fmt.Println(person3)

	//初始化的时候就取地址
	person4 := &Person{Age: 18, Name: "tom2"}
	person4.Name = "OK"
	fmt.Println(person4)

}

func A(person Person) {
	person.Age = 100
	person.Name = "A"
}

func B(person *Person) {
	person.Age = 111
	person.Name = "B"
}
