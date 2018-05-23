package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("Go Go Go!!!")
		//存数据操作
		c <- true
	}()

	//当main函数执行到这里后，就会阻塞在这里，等待c channel有数据被读取
	<-c
}
