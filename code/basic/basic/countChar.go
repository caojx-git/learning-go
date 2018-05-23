/**
 * 计算字符数
 */
package main

import "fmt"
import "unicode/utf8"

func main() {
	str := "asSASA ddd dsjkdsjs dk"
	fmt.Printf("String %s\nLength: %d, Runes: %d\n", str, len([]byte(str)), utf8.RuneCount([]byte(str)))
}
