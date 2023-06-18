package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"
)

type questions struct {
	Trait string `json:"trait"` // Which trait does this question measure?
	Text  string `json:"text"`
}

type quiz struct {
	Questions []questions `json:"questions"`
}

func handler(w http.ResponseWriter, r *http.Request, q *quiz) {
	switch r.Method {
	case "POST":
		handlePost(&w, r, q)
	case "GET":
		handleGet(&w, q)
	default:
		fmt.Print(r.Method)
		panic(" Unsupported request type!")
	}
}

func handleGet(w *http.ResponseWriter, q *quiz) {
	template := template.Must(template.ParseFiles("quiz.html"))
	err := template.Execute(*w, q)
	if err != nil {
		panic(err)
	}
}

func handlePost(w *http.ResponseWriter, r *http.Request, q *quiz) {
	r.ParseForm()
	traitScores := make(map[string]int)

	for index, response := range r.Form {
		responseValue, err := strconv.Atoi(response[0])
		if err != nil {
			panic(err)
		}
		indexValue, err := strconv.Atoi(index)
		if err != nil {
			panic(err)
		}

		trait := q.Questions[indexValue].Trait
		fmt.Println(index, response)
		traitScores[trait] += int(responseValue)
	}

	template := template.Must(template.ParseFiles("results.html"))
	err := template.Execute(*w, &traitScores)
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := os.ReadFile("questions.json")
	if err != nil {
		panic(err)
	}
	var quiz quiz
	json.Unmarshal(data, &quiz)

	// Create an anonymous function with the desired handler signature.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, &quiz)
	})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
