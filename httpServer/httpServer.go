package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func getDateTime(w http.ResponseWriter, req *http.Request) {

	currentTime := time.Now()
	response := currentTime
	json.NewEncoder(w).Encode(response)
}

func main() {

	http.HandleFunc("/datetime", getDateTime)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
