/**
 * goweb实现方式3
 */
package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler{},
		ReadTimeout: 5 * time.Second,
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/hello"] = sayHello
	mux["/sayBye"] = sayBye

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

type myHandler struct {
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "sayHello, this is version_2")
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "sayBye, this is version_2")
}
