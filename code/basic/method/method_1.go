/**
 * 方法
 * 通过显示说明receiver来实现与某个类型的组合
 */
package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{}
	a.Print()
	fmt.Println(a.Name)

	b := B{}
	b.print()
	fmt.Println(b.Name)
}

func (a *A) Print() {
	a.Name = "AA"
	fmt.Println("A")
}

func (b B) print() {
	b.Name = "BB"
	fmt.Println("B")
}
