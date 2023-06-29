package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "Hello World from go")
		if err != nil {
			log.Println("running with error:", err.Error())
		}
	})
	_ = http.ListenAndServe(":8080", nil) // http://localhost:8080
}
