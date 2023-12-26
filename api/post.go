package api

import (
	"encoding/json"
	"net/http"

	"mpxfactor.com/simple-server/data"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var employee data.Employee

		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		data.AddData(employee)
	} else {
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
}
