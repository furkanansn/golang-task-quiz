package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/furkanansn/golang-task-quiz/quiz"
)

type ServerConfig struct {
	Port int
}

type Server struct {
	Config *ServerConfig
	Router *mux.Router
	Quiz   *quiz.Quiz
	Scores *ScoreManager
}

func NewServer(cfg ServerConfig) (*Server, error) {
	server := &Server{
        Config: &cfg,
        Router: mux.NewRouter(),
        Quiz:   quiz.NewQuiz(),
        Scores: &ScoreManager{Scores: make(map[string]int)},
    }

    return server, nil
}

func (s *Server) SetupRoutes() {
	s.Router.HandleFunc("/api/questions", GetQuestionsHandler).Methods("GET")
	s.Router.HandleFunc("/api/submit-answers", SubmitAnswersHandler).Methods("POST")
	s.Router.HandleFunc("/api/user-score/{userID}", GetUserScoreHandler).Methods("GET")
	s.Router.HandleFunc("/api/compare-score/{userID}", CompareUserScoreHandler).Methods("GET")
}

func (s *Server) Start() {
	fmt.Printf("Server is running on port %d...\n", s.Config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.Config.Port), s.Router); err != nil {
		fmt.Printf("Error starting the server: %v", err)
	}
}
