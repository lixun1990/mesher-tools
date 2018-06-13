package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("hello")
		w.Write([]byte("hello"))

	})
	http.ListenAndServe(":9999", nil)
}
