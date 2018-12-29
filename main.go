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
	router.HandleFunc("/scan", ScanImage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func ScanImage(w http.ResponseWriter, r *http.Request) {
	var scanInstructions scanner.ScanInstructions
	_ = json.NewDecoder(r.Body).Decode(&scanInstructions)
	if err := scanner.Scan(scanInstructions); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(scanInstructions)
}
