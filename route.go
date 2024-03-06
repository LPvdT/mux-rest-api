package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{{Id: 1, Title: "Title 1", Text: "Lorem Upsum"}}
}

func getPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	result, err := json.Marshal(posts)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Cannot serialize posts array"}`))

		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(result)
}
