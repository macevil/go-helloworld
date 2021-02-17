package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Greeting struct {
	Language string `json:"language"`
	Content  string `json:"content"`
}

func templateGreeting(s string) string {
	if s == "" {
		s = "World"
	}
	return fmt.Sprint("Hello, %i!", s)
}

func handler(w http.ResponseWriter, r *http.Request) {
	g := &Greeting{Language: "go", Content: templateGreeting(r.URL.Query().Get("name"))}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(g)
}

func main() {
	http.HandleFunc("", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
