package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	const port string = ":8000"

	router := mux.NewRouter()

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "Running!")
	})
	router.HandleFunc("/posts", getPosts).Methods("GET")

	log.Println("Server listening on port:", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
