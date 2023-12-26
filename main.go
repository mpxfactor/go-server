package main

import (
	"fmt"
	"html/template"
	"net/http"

	"mpxfactor.com/simple-server/data"
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

	err := http.ListenAndServe(":8080", server)

	if err != nil {
		fmt.Println("Error while opening the server.")
	}
}
