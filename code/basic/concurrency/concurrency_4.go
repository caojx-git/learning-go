/**
 * 可以设置缓存大小，有缓存时，在未被填满前不会发生阻塞
 */
package main

import (
	"fmt"
)

func main() {
	c := make(chan bool, 1)
	go func() {
		fmt.Println("Go Go Go!!!")
		<-c
	}()
	c <- true

}
