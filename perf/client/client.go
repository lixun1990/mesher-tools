package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, _ := http.Get(os.Getenv("TARGET"))
		log.Println(resp)
		w.Write([]byte("hello"))

	})
	http.ListenAndServe(":9000", nil)
}
