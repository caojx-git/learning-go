/**
 * 基本类型也可添加方法
 */

package main

import (
	"fmt"
)

type TZ int

func main() {

	var a TZ
	//调用方式1
	a.Print()

	//调用方式2
	(*TZ).Print(&a)

}

func (a *TZ) Print() {
	fmt.Println("TZ")
}
