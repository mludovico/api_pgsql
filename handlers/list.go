package handlers

import (
	"api_pgsql/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request: ", r)
	rows, err := models.GetAll()
	if err != nil {
		log.Printf("Error getting all todos: %v", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Println("Response: ", rows)
	json.NewEncoder(w).Encode(rows)
}
