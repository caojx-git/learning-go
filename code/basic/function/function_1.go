/**
 * 定义函数
 */
package main

import (
	"fmt"
)

func main() {
	a, b := 1, 2
	F(a, b)
	fmt.Println(a, b)

	s1 := []int{1, 2, 3, 4, 5, 6}
	G(s1)
	fmt.Println(s1)
}

//函数A，无参数，无返回值
func A() {

}

/**
 *函数B，有参数，无返回值
 */
func B(a int, b string) {

}

/**
 * 函数C有参数，单返回值
 * 单个返回值可以省略括号
 */
func C(a int, b string) int {
	return 1
}

/**
 * 函数D有参数，多返回值,匿名返回值
 * 匿名返回必须需要return，切返回数据类型，对应函数返回值类型
 */
func D(a int, b string) (int, string, int) {
	c := 1
	return a, b, c
}

/*
 * 函数E有参数，多返回值,命名返回值
 * 命名返回 return是可以省略对应的返回名
 */
func E(a int, b int) (d, e, f int) {
	d, e, f = a, b, 3
	return
}

/**
 * 不定长边参数，不定长变参数必须作为函数中的最后一个参数
 * 变参相当于 slice, 但传递的是值拷贝
 */
func F(a ...int) {
	a[0] = 2
	a[1] = 2
	fmt.Println(a)
}

/**
 *如果直接传递一个slice，则传递的是指针地址
 */
func G(s []int) {
	s[0] = 3
	s[1] = 3
	fmt.Println(s)
}
