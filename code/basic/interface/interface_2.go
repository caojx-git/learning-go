/*
接口可以嵌入到其他接口，或嵌入到结构中
*/
package main

import (
	"fmt"
)

/**
 * 接口可以匿名嵌入其他接口，或嵌入到结构中
 */
type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connect:", pc.name)
}

func main() {
	var usb USB
	usb = PhoneConnecter{name: "PhoneConnecter"}
	usb.Connect()
	Discconect(usb)

	/**
	* 接口转换，只能用拥有超集的接口转化为它子集的接口，如上边的USB可以转为Connecter接口，但是Connecter
	* 接口不能转为USB接口
	 */
	pc := PhoneConnecter{"PhoneConnecter"}
	var a Connecter
	a = Connecter(pc)
	a.Connect()
	//a.Name undefined (type Connecter has no field or method Name)
	//a.Name()
}

/*func Discconect(usb USB) {
	if pc, ok :=usb.(PhoneConnecter); ok {
		fmt.Println("Disconnected:",pc.name)
		return
	}
	fmt.Println("Unknown decive.")
}
*/

/*
所有的类型都实现了空接口
*/
func Discconect(usb interface{}) {

	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknown decive.")
	}

}
