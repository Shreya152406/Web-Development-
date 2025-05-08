package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

const baseURL = "https://studentvotingsystem-91633-default-rtdb.asia-southeast1.firebasedatabase.app"
const votesURL = baseURL + "/votes.json"
const voteCountsURL = baseURL + "/voteCounts.json"

type VoteData struct {
	Enrollment string `json:"enrollment"`
	Candidate  string `json:"candidate"`
}

type PageData struct {
	VoteCounts map[string]int
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		fullname := r.FormValue("fullname")
		enrollment := r.FormValue("enrollment")
		contact := r.FormValue("contact")
		branch := r.FormValue("branch")

		log.Println("Received Student Data:", fullname, enrollment, contact, branch)

		tmpl, err := template.ParseFiles("templates/register.html")
		if err != nil {
			log.Println("Error loading register.html:", err)
			http.Error(w, "Error loading page", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, map[string]string{
			"Enrollment": enrollment,
		})
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func voteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handleVoteSubmission(w, r)
		return
	}

	// Handle GET request (show vote counts)
	resp, err := http.Get(voteCountsURL)
	if err != nil {
		log.Println("Error getting vote counts:", err)
		http.Error(w, "Error loading counts", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	voteCounts := make(map[string]int)
	if string(body) != "null" {
		json.Unmarshal(body, &voteCounts)
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Println("Template error:", err)
		http.Error(w, "Template load error", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, PageData{VoteCounts: voteCounts})
}

func handleVoteSubmission(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	enrollment := r.FormValue("enrollment")
	candidate := r.FormValue("candidate")

	// Check for existing vote
	resp, err := http.Get(votesURL)
	if err != nil {
		log.Println("Error fetching votes:", err)
		http.Error(w, "Database access error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var allVotes map[string]VoteData
	if string(body) != "null" {
		json.Unmarshal(body, &allVotes)
	}

	for _, vote := range allVotes {
		if vote.Enrollment == enrollment {
			http.Redirect(w, r, "/vote?voted=2", http.StatusSeeOther)
			return
		}
	}

	// Submit new vote
	vote := VoteData{Enrollment: enrollment, Candidate: candidate}
	voteJSON, _ := json.Marshal(vote)
	_, err = http.Post(votesURL, "application/json", bytes.NewBuffer(voteJSON))
	if err != nil {
		log.Println("Error posting new vote:", err)
		http.Error(w, "Vote submission error", http.StatusInternalServerError)
		return
	}

	// Update vote count
	updateVoteCount(candidate)

	http.Redirect(w, r, "/vote?voted=1", http.StatusSeeOther)
}

func updateVoteCount(candidate string) {
	// Get current vote counts
	resp, err := http.Get(voteCountsURL)
	if err != nil {
		log.Println("Error getting vote counts:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	voteCounts := make(map[string]int)
	if string(body) != "null" {
		json.Unmarshal(body, &voteCounts)
	}

	voteCounts[candidate]++

	updatedVoteCountJSON, _ := json.Marshal(voteCounts)
	req, _ := http.NewRequest(http.MethodPut, voteCountsURL, bytes.NewBuffer(updatedVoteCountJSON))
	req.Header.Set("Content-Type", "application/json")
	resp2, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Failed to update vote counts:", err)
		return
	}
	defer resp2.Body.Close()
}

func main() {
	http.HandleFunc("/", voteHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/vote", voteHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Server started at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed:", err)
	}
}
