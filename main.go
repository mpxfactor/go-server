package main

import (
	"fmt"
	"net/http"
)

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	message := []byte("First http get request in go.")
	w.Write(message)
}

func main() {
	server := http.NewServeMux() //serve multiplexer
	server.HandleFunc("/get", handleGetRequest)

	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)

	err := http.ListenAndServe(":8080", server)

	if err != nil {
		fmt.Println("Error while opening the server.")
	}
}
