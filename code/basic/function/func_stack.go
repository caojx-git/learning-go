/**
 * 模拟栈操作
 */
package main

import "fmt"

type stack struct {
	i    int
	data [10]int
}

func (s *stack) push(k int) {
	if s.i+1 > 9 {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *stack) pop() int {
	s.i--
	return s.data[s.i]
}

/*
func (s stack) String() string {
	var str string
	for i := 0; i <= s.i; i++ {
		str = str + "["+strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i]) + "]"
	}
	return str
}
*/

func main() {
	var s stack
	s.push(25)
	s.push(14)
	fmt.Printf("stack %v\n", s)

	//str := s.String();
	//fmt.Printf(str)
}
