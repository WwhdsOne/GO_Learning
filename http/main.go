package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello go"))
	})

	// 刷新token

	http.ListenAndServe(":8848", nil)

	//
}
