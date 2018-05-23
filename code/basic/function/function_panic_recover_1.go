/**
 * Go没有异常机制，但有panic/ercover模式来处理错误
 * Panic可以在任何地方引发，但recover只有defer调用的函数中有效
 */

package main

import (
	"fmt"
)

func main() {
	A()
	//B发生panic程序终止，不会执行C
	B()
	C()
}

func A() {
	fmt.Println("func A")
}

func B() {
	//在panic之前注册defer函数，从panic中恢复
	defer func() {
		//判断是否发生panic
		if err := recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()

	//发生panic
	panic("Panic in B")

}

func C() {
	fmt.Println("Func c")
}
