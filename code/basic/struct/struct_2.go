/**
 * 匿名结构
 */
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
	//嵌入匿名结构
	Contact struct {
		Phone, City string
	}
}

//匿名字段，初始化值的顺序要与字段类型顺序一致
type test struct {
	string
	int
}

func main() {
	//匿名结构
	a := struct {
		Name string
		Age  int
	}{
		Name: "joe",
		Age:  19,
	}

	fmt.Println(a)

	//嵌套结构
	b := Person{Name: "joe", Age: 19}
	b.Contact.Phone = "123456"
	b.Contact.City = "beijin"

	fmt.Println(b)

	//匿名字段
	c := test{"a", 1}
	fmt.Println(c)

	d := Person{Name: "joe", Age: 19}

}
