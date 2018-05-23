/**
 * 变参
 */
package main

import "fmt"

func prtthem(numbers ...int) {
	for _, d := range numbers {
		fmt.Printf("%d\n", d)
	}
}

func main() {
	prtthem(1, 2, 3, 4, 5)
}
