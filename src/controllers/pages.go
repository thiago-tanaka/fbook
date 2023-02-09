package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/models"
	"webapp/src/requests"
	"webapp/src/responses"
	"webapp/src/utils"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func LoadCreateAccountPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "create-account.html", nil)
}

func LoadHomePage(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("%s/posts", config.APIURL)
	response, err := requests.MakeRequestWithAuth(r, http.MethodGet, url, nil)

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

	fmt.Println(response.StatusCode)

	var posts []models.Post

	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ApiError{
			Message: err.Error(),
		})
		return
	}

	utils.ExecuteTemplate(w, "home.html", posts)
}
