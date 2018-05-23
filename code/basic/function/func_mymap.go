/**
 * 定义map映射
 */
package main

import "fmt"

func mymap(f func(int) int, l []int) []int {
	j := make([]int, len(l))
	for k, v := range l {
		j[k] = f(v)
	}
	return j
}

func main() {
	m := []int{1, 2, 3, 4}
	f := func(i int) int {
		return i * i
	}
	fmt.Printf("%v", (mymap(f, m)))
}
