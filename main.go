package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Course struct {
	Team     string
	Champion string
	Year     int
}

func (c Course) getFullInfo() string {
	return fmt.Sprintf("Team: %s, Chamnpion: %s, Year: %d", c.Team, c.Champion, c.Year)
}

func home(w http.ResponseWriter, r *http.Request) {
	course := Course{
		Team:     "Palmeiras",
		Champion: "Libertadores",
		Year:     2023,
	}

	// Marshal the Course object to JSON
	jsonBytes, err := json.Marshal(course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonBytes)
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
