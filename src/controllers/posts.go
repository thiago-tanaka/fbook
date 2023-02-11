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
	//post := struct {
	//	Title   string `json:"title"`
	//	Content string `json:"content"`
	//}{
	//	Title:   r.FormValue("title"),
	//	Content: r.FormValue("content"),
	//}
	post, err := json.Marshal(map[string]string{
		"title":   r.FormValue("title"),
		"content": r.FormValue("content"),
	})

	//postJSON, err := json.Marshal(post)
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
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

	responses.JSON(w, response.StatusCode, nil)
}

func likePost(w http.ResponseWriter, r *http.Request, action string) {
	params := mux.Vars(r)
	postID, err := strconv.ParseUint(params["postId"], 10, 64)
	if err != nil {
		responses.JSON(w, http.StatusBadRequest, responses.ApiError{
			Message: "invalid post ID",
		})
		return
	}

	url := fmt.Sprintf("%s/posts/%d/%s", config.APIURL, postID, action)
	response, err := requests.MakeRequestWithAuth(r, http.MethodPost, url, nil)
	if err != nil {
		responses.JSON(w, http.StatusInternalServerError, responses.ApiError{
			Message: err.Error(),
		})
		return
	}
	defer response.Body.Close()

	responses.JSON(w, response.StatusCode, nil)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	likePost(w, r, "like")
}

func DislikePost(w http.ResponseWriter, r *http.Request) {
	likePost(w, r, "dislike")
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
