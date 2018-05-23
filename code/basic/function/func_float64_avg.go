/**
 * 求平均数
 */
package main

import "fmt"

func average(s []float64) float64 {
	sum := 0.0
	avg := 0.0
	if len(s) == 0 {
		avg = 0.0
	}
	for _, v := range s {
		sum += v
	}
	avg = sum / float64(len(s))
	return avg
}

func main() {
	var a = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	avg := average(a)
	fmt.Printf("%f\n", avg)
}
