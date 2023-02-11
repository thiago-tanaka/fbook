package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
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

	var posts []models.Post

	if err = json.NewDecoder(response.Body).Decode(&posts); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ApiError{
			Message: err.Error(),
		})
		return
	}

	cookie, _ := cookies.Read(r)
	userID, _ := strconv.ParseUint(cookie["ID"], 10, 64)

	utils.ExecuteTemplate(w, "home.html", struct {
		Posts  []models.Post
		UserID uint64
	}{
		Posts:  posts,
		UserID: userID,
	})
}

func LoadPostEditPage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID := params["postId"]

	url := fmt.Sprintf("%s/posts/%s", config.APIURL, postID)

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

	var post models.Post

	if err = json.NewDecoder(response.Body).Decode(&post); err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ApiError{
			Message: err.Error(),
		})
		return
	}

	utils.ExecuteTemplate(w, "edit-post.html", post)
}
