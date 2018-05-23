/**
 * 冒泡排序
 */
package main

import "fmt"

func bubblesort(n []int) {
	for i := 0; i < len(n)-1; i++ {
		for j := i + 1; j < len(n)-1; j++ {
			if n[i] > n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
}

func main() {
	n := []int{5, -1, 0, 12, 3, 5}
	fmt.Printf("unsorted %v\n", n)
	bubblesort(n)
	fmt.Printf("sorted %v\n", n)
}
