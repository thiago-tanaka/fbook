package controllers

import (
	"net/http"
	"webapp/src/utils"
)

func LoadLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}