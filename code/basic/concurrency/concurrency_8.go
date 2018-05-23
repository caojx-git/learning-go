package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	//同时有多个可用的channel时安随机顺序处理
	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}
}
