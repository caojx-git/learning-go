package main
import "fmt"

func main() {
	str := "A"
	for i := 1; i < 100; i++ {
		fmt.Printf("%s\n", str)
		str = str + "A"
	}
}
