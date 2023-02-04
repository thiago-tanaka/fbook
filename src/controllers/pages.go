package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

func LoadCreateAccountPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "create-account.html", nil)
}
