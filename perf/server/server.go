package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hello")
		w.Write([]byte("hello"))

	})
	http.ListenAndServe(":9999", nil)
}
