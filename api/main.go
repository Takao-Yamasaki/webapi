package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		lang := r.URL.Query().Get("lang")
		if lang == "en" {
			fmt.Fprintf(w, "HELLO")
		} else {
			fmt.Fprintf(w, "こんにちは")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
