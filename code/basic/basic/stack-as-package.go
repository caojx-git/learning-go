package stack

//结构Stack
type Stack struct {
	i    int
	data [10]int
}

//push 将元素压入栈中
func (s *Stack) Push(k int) {
	s.data[s.i] = k
	s.i++
}

//pop 从栈中弹出一个元素
func (s *Stack) Pop() int {
	s.i--
	return s.data[s.i]
}
