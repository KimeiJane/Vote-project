package main

import (
	"html/template"
	"log"
	"net/http"
	"sync"
)

// In-memory vote counters
var votes = map[string]int{
	"Go":        0,
	"Python":    0,
	"JavaScript": 0,
	"Java":      0,
}

var mu sync.Mutex // mutex to protect concurrent access

func main() {
	fs := http.FileServer(http.Dir("../Frontend"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", votePage)
	http.HandleFunc("/vote", voteHandler)
	http.HandleFunc("/results", resultsPage)
	http.HandleFunc("/reset", resetVotes)

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func votePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../Frontend/vote.html")
}

func resultsPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("../Frontend/results.html")
	if err != nil {
		http.Error(w, "Template parsing error", http.StatusInternalServerError)
		return
	}
	mu.Lock()
	defer mu.Unlock()
	tmpl.Execute(w, votes)
}

func voteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		choice := r.FormValue("language")
		mu.Lock()
		if _, exists := votes[choice]; exists {
			votes[choice]++
		}
		mu.Unlock()
		http.Redirect(w, r, "/results", http.StatusSeeOther)
	}
}

func resetVotes(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	for k := range votes {
		votes[k] = 0
	}
	mu.Unlock()
	http.Redirect(w, r, "/", http.StatusSeeOther)
}