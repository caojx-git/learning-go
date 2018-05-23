package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", Cookie)
	http.HandleFunc("/2", Cookie2)
	http.ListenAndServe(":8080", nil)
}

/**
 * 浏览器需要第二次才可以将cookie读取出来
 */
func Cookie(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "myCookie",
		Value:  "hello",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}

	http.SetCookie(w, ck)

	ck2, err := r.Cookie("myCookie")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, ck2.Value)
}

/**
 * 浏览器可以一次读取出Cooke,但是Cookie内容值度空格兼容性不好
 */
func Cookie2(w http.ResponseWriter, r *http.Request) {
	ck := &http.Cookie{
		Name:   "myCookie2",
		Value:  "helloworld",
		Path:   "/",
		Domain: "localhost",
		MaxAge: 120,
	}

	w.Header().Set("Set-Cookie", ck.String())

	ck2, err := r.Cookie("myCookie2")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, ck2.Value)
}
