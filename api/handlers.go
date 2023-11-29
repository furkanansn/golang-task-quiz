package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var server *Server

func InitServer(s *Server) {
	server = s
}

type ScoreManager struct {
	Scores map[string]int
}

func (sm *ScoreManager) SetScore(userID string, score int) {
	sm.Scores[userID] = score
}

func (sm *ScoreManager) GetScore(userID string) int {
	return sm.Scores[userID]
}

func (sm *ScoreManager) GetPercentageBetterThanUser(userScore int) int {
	var betterCount int
	for _, score := range sm.Scores {
		if score > userScore {
			betterCount++
		}
	}

	totalUsers := len(sm.Scores)
	if totalUsers == 0 {
		return 0
	}

	return (betterCount * 100) / totalUsers
}

func GetQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	questions := server.Quiz.Questions
	err := json.NewEncoder(w).Encode(questions)
	if err != nil {
		log.Printf("Error encoding JSON response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func SubmitAnswersHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody struct {
		UserID  string         `json:"userID"`
		Answers map[string]int `json:"answers"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	userID := requestBody.UserID
	userAnswers := requestBody.Answers
	if userID == "" || len(userAnswers) == 0 {
		http.Error(w, "Invalid user ID or answers", http.StatusBadRequest)
		return
	}

	var answers []int
	for i := 1; i <= len(userAnswers); i++ {
		answer, ok := userAnswers["q"+strconv.Itoa(i)]
		if !ok {
			http.Error(w, fmt.Sprintf("Missing answer for question %d", i), http.StatusBadRequest)
			return
		}
		answers = append(answers, answer)
	}

	score := server.Quiz.Score(answers)

	server.Scores.SetScore(userID, score)

	w.Write([]byte(fmt.Sprintf("Your score: %d", score)))
}

func GetUserScoreHandler(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]
	score := server.Scores.GetScore(userID)
	w.Write([]byte(fmt.Sprintf("User %s's score: %d", userID, score)))
}

func CompareUserScoreHandler(w http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]
	userScore := server.Scores.GetScore(userID)
	percentage := server.Scores.GetPercentageBetterThanUser(userScore)

	w.Write([]byte(fmt.Sprintf("You were better than %d%% of all quizzers", percentage)))
}
