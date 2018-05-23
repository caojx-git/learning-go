/**
 * 斐波
 */

package main

import "fmt"

func fibnacci(value int) []int {
	var x = make([]int, value)
	x[0], x[1] = 1, 1
	for n := 2; n < value; n++ {
		x[n] = x[n-1] + x[n-2]
	}
	return x
}

func main() {
	for _, term := range fibnacci(10) {
		fmt.Printf("%v ", term)
	}
}
