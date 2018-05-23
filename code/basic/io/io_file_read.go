package main 

import "os"

func main() {
	buf := make([]byte, 1024)
	//打开文件，os.Open 返回一个实现了 io.Reader 和 io.Writer 的 *os.File;
	f, _ := os.Open("/etc/passwd")
	//确保关闭了 f;
	defer f.Close()

	for {
		//一次读取 1024 字节
		n, _ := f.Read(buf)
		//到达文件末尾;
		if n == 0 {
			break
		}
		//将内容写入 os.Stdout
		os.Stdout.Write(buf[:n])
	}
}