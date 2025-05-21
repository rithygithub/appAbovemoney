package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func aspirationsHandler(w http.ResponseWriter, r *http.Request) {
	// Load your aspirations or use your existing code here
	content, err := ioutil.ReadFile("aspirations.json")
	if err != nil {
		http.Error(w, "Failed to read aspirations", http.StatusInternalServerError)
		return
	}

	// Print the same output as before, or render HTML/JSON as you prefer
	fmt.Fprintf(w, "Aspirations for Day 141, Hour 04:\n")
	fmt.Fprintf(w, "%s", content)
}

func main() {
	http.HandleFunc("/", aspirationsHandler)
	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}