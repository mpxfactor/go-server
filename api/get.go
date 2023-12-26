package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"mpxfactor.com/simple-server/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// /api/employees?id=1
	id := r.URL.Query()["id"]
	if id != nil {
		finalId, err := strconv.Atoi(id[0])

		if err == nil && 0 < finalId && finalId < len(data.GetAllData())+1 {
			json.NewEncoder(w).Encode(data.GetAllData()[finalId-1])
		} else {
			http.Error(w, "Ivalid ID", http.StatusBadRequest)
		}
	} else {
		json.NewEncoder(w).Encode(data.GetAllData())
	}

}
