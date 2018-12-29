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
		w.Write([]byte(err.Error()))
		return
	}
	jsonData := map[string]string{"path": result}
	json.NewEncoder(w).Encode(jsonData)
}
