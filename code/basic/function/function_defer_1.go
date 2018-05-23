/**
 * defer
 */
package main

import (
	"fmt"
)

func main() {

	//1.执行方式类似于其他语言中的析构函数，在函数体执行结速后，按照调用顺序的相反顺序逐个执行
	//fmt.Println("a")
	//defer fmt.Println("b")
	//defer fmt.Println("c")

	//2.
	/*
		for i := 0; i < 3; i++ {
			defer fmt.Println(i)
		}
	*/

	//3.
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}

}
