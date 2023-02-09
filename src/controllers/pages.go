package controllers

import (
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/requests"
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

	response, _ := requests.MakeRequestWithAuth(r, http.MethodGet, url, nil)

	fmt.Println(response.StatusCode)

	utils.ExecuteTemplate(w, "home.html", nil)
}
