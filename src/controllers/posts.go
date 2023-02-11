package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/requests"
	"webapp/src/responses"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ApiError{
			Message: err.Error(),
		})
		return
	}

	url := fmt.Sprintf("%s/posts", config.APIURL)

	response, err := requests.MakeRequestWithAuth(r, http.MethodPost, url, bytes.NewBuffer(post))

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

func LikePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)

	url := fmt.Sprintf("%s/posts/%d/like", config.APIURL, postID)

	response, err := requests.MakeRequestWithAuth(r, http.MethodPost, url, nil)

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

func DeletePost(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)

	url := fmt.Sprintf("%s/posts/%d", config.APIURL, postID)

	response, err := requests.MakeRequestWithAuth(r, http.MethodDelete, url, nil)

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
