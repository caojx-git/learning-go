/**
 * 匿名函数与闭包
 */

package main

import (
	"fmt"
)

func main() {
	/**
	 * 匿名函数
	 */
	a := func() {
		fmt.Println("Func A")
	}
	a()

	/**
	 * 闭包
	 */
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}

/**
 * 闭包
 */
func closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
