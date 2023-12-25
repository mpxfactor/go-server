package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("First http request in go."))
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error while opening the server.")
	}
}
