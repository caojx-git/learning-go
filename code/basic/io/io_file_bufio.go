package main 

import (
	"os"
	"bufio"
)

func main() {
	buf := make([]byte, 1024)
	//打开文件;
	f, _ := os.Open("/etc/passwd")
	defer f.Close()

	//转换 f 为有缓冲的 Reader。NewReader 需要一个 io.Reader，因此或许你认为 这会出错。但其实不会。任何有 Read() 函数就实现了这个接口。
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	for {
		//从 Reader 读取，而向 Writer 写入，然后向屏幕输出文件。
		n, _ := r.Read(buf)
		if n == 0 {
			break
		}
		w.Write(buf[:n])
	}
}