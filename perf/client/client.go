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
		b := make([]byte, 0)
		_, _ = resp.Body.Read(b)
		w.Write(b)

	})
	http.ListenAndServe(":9000", nil)
}
