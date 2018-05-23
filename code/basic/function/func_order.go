/**
 * 排序
 */
package main

import "fmt"

func order(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func main() {
	a, b := 1, 2
	c, d := order(a, b)
	fmt.Printf("before: %d, %d\n", a, b)
	fmt.Printf("after %d, %d\n", c, d)
}
