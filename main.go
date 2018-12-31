package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"scantastic/scanner"
	"scantastic/thumbify"
)

type RequestParams struct {
	scanner.ScanInstructions
	Thumbnail bool `json: thumbnail`
}

// our main function
func main() {
	router := mux.NewRouter()
	scanner.Init()
	defer scanner.End()
	thumbify.Start()
	defer thumbify.End()
	router.HandleFunc("/scan", scanImage).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func scanImage(w http.ResponseWriter, r *http.Request) {
	var requestParams RequestParams
	_ = json.NewDecoder(r.Body).Decode(&requestParams)
	log.Println(requestParams)
	filename, err := scanner.Scan(requestParams.ScanInstructions)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonData := map[string]string{"error": err.Error()}
		_ = json.NewEncoder(w).Encode(jsonData)
		return
	}
	var thumb string
	if requestParams.Thumbnail {
		bytes, err := thumbify.ThisImage(filename)
		if err != nil {
			thumb = "ERROR"
		} else {
			thumb = string(bytes)
		}
	}
	jsonData := map[string]string{"filename": filename, "thumbnail": thumb}
	_ = json.NewEncoder(w).Encode(jsonData)
}
