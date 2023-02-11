package responses

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ApiError struct {
	Message string `json:"message"`
}

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if status != http.StatusNoContent {
		fmt.Println("data: ", data)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Fatal(err)
		}
	}

}

func HandleError(w http.ResponseWriter, r *http.Response) {
	var err ApiError

	json.NewDecoder(r.Body).Decode(&err)
	JSON(w, r.StatusCode, err)
}
