/**
 * goweb实现方式1
 */
package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	//设置路由
	http.HandleFunc("/", sayHello)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world, this is version_1")
}
