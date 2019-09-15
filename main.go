package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test title", Desc: "description", Content: "content od the whole thing"},
		Article{Title: "Test title 2", Desc: "description 2", Content: "content od the whole thing 2"},
	}

	fmt.Printf("Endpoint hit: all articles")

	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage endpoint Hit")
}

func hotels(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hotels endpoint Hit")

	json.NewEncoder(w).Encode("wip")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	http.HandleFunc("/hotels", hotels)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
