/*
只有当接口存储的类型和对象都为nil时，接口才等于nil
*/
package main

import (
	"fmt"
)

func main() {
	var a interface{}
	fmt.Println(a == nil)

	var p *int = nil

	a = p
	fmt.Println(a == nil)
}
