package main

import (
	"fmt"
	"html/template"
	"net/http"

	"mpxfactor.com/simple-server/data"
)

const (
	PORT = "8081"
)

func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.tmpl")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error."))
		return
	}

	html.Execute(w, data.GetAllData())
}

func main() {
	server := http.NewServeMux() //serve multiplexer
	server.HandleFunc("/employees", handleTemplate)

	fmt.Printf("Server running on : http://localhost:%s", PORT)
	err := http.ListenAndServe(":"+PORT, server)

	if err == nil {
		fmt.Println("Error while opening the server.")
		fmt.Println(err)
	}
}
