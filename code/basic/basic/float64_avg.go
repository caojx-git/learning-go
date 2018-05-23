/**
 * 求平均值
 */
package main

import "fmt"

func main() {
	a := [10]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s := a[0 : len(a)-1]
	average(s)
}

func average(xs []float64) {
	sum := 0.0
	avg := 0.0
	switch len(xs) {
	case 0:
		avg = 0
	default:
		for _, v := range xs {
			sum += v
		}
		avg = sum / float64(len(xs))
	}
	fmt.Println(avg)
}
