/**
 * 结构之间的赋值与比较
 */
package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	a := Person{Name: "joe", Age: 19}

	var b Person

	b = a
	fmt.Println(b)

	c := Person{Name: "joe", Age: 19}

	//支持== 与!= 比较运算符， 但不支持 > 或 <
	fmt.Println(a == c)
}
