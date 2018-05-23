/*
只要某个类型拥有该接口的所有方法签名，即实现了改接口，无需现实实现了那个接口，这个称为Structural Typing
*/
package main

import (
	"fmt"
)

/**
 * 接口只有方法声明，没有实现，没有数据字段
 */
type USB interface {
	Name() string
	Connect()
}

type PhoneConnecter struct {
	name string
}

/*
只要某个类型拥有该接口的所有方法签名，即实现了改接口，无需现实实现了那个接口，这个称为Structural Typing
PhoneConnecter 方法实现 USB 接口中的Name方法
*/
func (pc PhoneConnecter) Name() string {
	return pc.name
}

/*
只要某个类型拥有该接口的所有方法签名，即实现了改接口，无需现实实现了那个接口，这个称为Structural Typing
PhoneConnecter 方法实现 USB 接口中的Name方法
*/
func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func main() {
	var usb USB
	usb = PhoneConnecter{name: "PhoneConnecter"}

	//usb := PhoneConnecter{name: "PhoneConnecter"}

	// 输出：Connect: PhoneConnecter
	usb.Connect()
	Discconect(usb)
}

func Discconect(usb USB) {
	fmt.Println("Disconnected")
}
