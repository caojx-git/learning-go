package main
import "fmt"

func main() {
	s := "foobar"
	s2 := reverse(s)
	fmt.Println(s2)
}

func reverse(str string) string {
	rs := []rune(str)
	len := len(rs)
	var tt []rune

	tt = make([]rune, 0)
	for i := 0; i < len; i++ {
		tt = append(tt, rs[len-i-1])
	}
	return string(tt)
}
