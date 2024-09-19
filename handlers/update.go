package handlers

import (
	"api_pgsql/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request: ", r)
	fmt.Printf("Request: %v\n", r)
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	var todo models.Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rowsAffected, err := models.Update(int64(id), todo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rowsAffected > 1 {
		log.Printf("Error. %v rows affected.", rowsAffected)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Todo updated successfully",
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Println("Response: ", resp)
	json.NewEncoder(w).Encode(resp)
}
