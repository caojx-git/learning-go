package main 

import ("io/ioutil"; "net/http"; "fmt")

func main() {
	//使用 http 的 Get 获取 html;
	r, err := http.Get("https://www.baidu.com/robots.txt")

	//错误处理;
	if err != nil {
		fmt.Printf("%s\n", err.String())
		return
	}

	//将整个内容读入 b;
	b, err := ioutil.ReadAll(r.Body)
	r.Body.Close()

	//如果一切 OK 的话，打印内容。
	if err == nil {
		fmt.Printf("%s", stirng(b))
	}
}