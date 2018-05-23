/**
 * goweb实现方式2
 */
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	//注册路由
	mux.Handle("/", &myHandler{})

	//注册函数
	mux.HandleFunc("/hello", sayHello)

	//静态文件处理
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(wd))))

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct {
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL:"+r.URL.String())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world, this is version_2")
}
