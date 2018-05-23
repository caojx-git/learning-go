/*
将对象赋值给接口时，会发生拷贝，而接口内部存储的是指向这个复制品的指针，即无法修改复制品的状态，也无法获取指针。
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
	pc := PhoneConnecter{"PhoneConnecter"}
	var a Connecter
	a = Connecter(pc)
	a.Connect()

	pc.name = "pc"
	a.Connect()

	//Connect: PhoneConnecter
	//Connect: PhoneConnecter
}

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
