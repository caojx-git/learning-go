package main
import "fmt"

func main() {
	s := "asSASA ddd dsjkdsjs dk"
	r := []rune(s)
	copy(r[3:7],[]rune("abc"))
	fmt.Println("before %s\n", s)
	fmt.Println("after %s\n", string(r))
}
