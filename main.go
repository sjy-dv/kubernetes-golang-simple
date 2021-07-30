package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloKuber(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"result": "hello kuber",
	})
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/helloapi", HelloKuber).Methods("GET")

	http.ListenAndServe(":8081", r)
}
