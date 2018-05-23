/**
 * 可以使用for range来迭代不断操作channel
 */
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
		//关闭之后，for循环会停止等待,如果不关闭，则下边的for循环会无限等待
		close(c)
	}()

	//通过无限循环，等待channel被输出
	for v := range c {
		fmt.Println(v)
	}
}
