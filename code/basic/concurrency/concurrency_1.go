package main

import (
	"fmt"
	"time"
)

func main() {
	//由于使用了多线程机制，方法还没有调用完，主线程main方法就已经执行完了，下边啥也不会输出
	go Go()
	//time.Sleep(2 * time.Second)
}

func Go() {
	fmt.Println("Go Go Go!!!")
}
