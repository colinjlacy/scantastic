package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"scantastic/scanner"
)

// our main function
func main() {
	router := mux.NewRouter()
	scanner.Init()
	defer scanner.End()
	router.HandleFunc("/scan", scanImage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func scanImage(w http.ResponseWriter, r *http.Request) {
	var scanInstructions scanner.ScanInstructions
	_ = json.NewDecoder(r.Body).Decode(&scanInstructions)
	result, err := scanner.Scan(scanInstructions)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonData := map[string]string{"error": err.Error()}
		jsonValue, _ := json.Marshal(jsonData)
		json.NewEncoder(w).Encode(jsonValue)
		return
	}
	jsonData := map[string]string{"path": result}
	jsonValue, _ := json.Marshal(jsonData)
	json.NewEncoder(w).Encode(jsonValue)
}
