package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var baseURL = "http://localhost:8080"

func main() {
	var rootCmd = &cobra.Command{Use: "quiz-cli"}

	var questionsCmd = &cobra.Command{
		Use:   "questions",
		Short: "Get quiz questions",
		Run: func(cmd *cobra.Command, args []string) {
			getQuestions()
		},
	}

	var submitCmd = &cobra.Command{
		Use:   "submit",
		Short: "Submit quiz answers",
		Run: func(cmd *cobra.Command, args []string) {
			userID := args[0]
			submitAnswers(userID)
		},
	}

	var userScoreCmd = &cobra.Command{
		Use:   "user-score",
		Short: "Get user score",
		Run: func(cmd *cobra.Command, args []string) {
			userID := args[0]
			getUserScore(userID)
		},
	}

	var compareScoreCmd = &cobra.Command{
		Use:   "compare-score",
		Short: "Compare user score",
		Run: func(cmd *cobra.Command, args []string) {
			userID := args[0]
			compareUserScore(userID)
		},
	}

	rootCmd.AddCommand(questionsCmd, submitCmd, userScoreCmd, compareScoreCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getQuestions() {
	url := fmt.Sprintf("%s/api/questions", baseURL)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting questions: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Printf("Quiz Questions: %s\n", body)
}

func submitAnswers(userID string) {
	answers := map[string]int{"q1": 0, "q2": 2, "q3": 1}

	payload := map[string]interface{}{
		"UserID":  userID,
		"answers": answers,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		return
	}

	url := fmt.Sprintf("%s/api/submit-answers", baseURL)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error submitting answers: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Printf("Server response: %s\n", body)
}

func getUserScore(userID string) {
	url := fmt.Sprintf("%s/api/user-score/%s", baseURL, userID)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error getting user score: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Printf("User Score: %s\n", body)
}

func compareUserScore(userID string) {
	url := fmt.Sprintf("%s/api/compare-score/%s", baseURL, userID)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error comparing user score: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	fmt.Printf("Comparison Result: %s\n", body)
}
