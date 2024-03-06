package main

import (
	"encoding/json"
	"net/http"
)

const port string = ":8000"

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

// The function init generates placeholder data
func init() {
	posts = []Post{{Id: 1, Title: "Title 1", Text: "Lorem Upsum"}}
}

// The function getPosts retrieves and returns a list of posts in JSON format
// over HTTP.
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

// The function addPost reads a JSON request, creates a new post with an
// incremented ID, adds it to an array, and returns the serialized post as a
// response.
func addPost(res http.ResponseWriter, req *http.Request) {
	var post Post

	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Cannot deserialize request"}`))
		return
	}

	// Increment ID and add to array
	post.Id = len(posts) + 1
	posts = append(posts, post)

	res.WriteHeader(http.StatusOK)

	// Serialize and write
	result, err := json.Marshal(post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Cannot serialize new post"}`))
		return
	}

	res.Write(result)
}
