package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/responses"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	user, err := json.Marshal(map[string]string{
		"name":     r.FormValue("name"),
		"email":    r.FormValue("email"),
		"nick":     r.FormValue("nick"),
		"password": r.FormValue("password"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ApiError{
			Message: err.Error(),
		})
		return
	}

	response, err := http.Post("http://localhost:5000/users", "application/json", bytes.NewBuffer(user))

	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ApiError{
			Message: err.Error(),
		})
		return
	}

	defer response.Body.Close()

	if response.StatusCode >= 400 {
		responses.HandleError(w, response)
		return
	}

	responses.JSON(w, response.StatusCode, nil)
}
