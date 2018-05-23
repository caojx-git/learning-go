package main

import (
	"fmt"
	"runtime"
)

func main() {
	//使用多个CPU
	runtime.GOMAXPROCS(runtime.NumCPU())
	//使用channel实现，完成后才结束主线程
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go Go(c, i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}

func Go(c chan bool, index int) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	c <- true
}
