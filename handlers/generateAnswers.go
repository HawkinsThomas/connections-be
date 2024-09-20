package handlers

import (
    "encoding/json"
    "net/http"
		"connections/sources/openai"
		"log"
)

type GenerateAnswersResponse struct {
    Clue    string   `json:"clue"`
    Answers []string `json:"answers"`
}

func GenerateAnswersHandler(w http.ResponseWriter, r *http.Request) {
		clue := r.URL.Query().Get("clue")

		if clue == "" {
			http.Error(w, "Clue is required", http.StatusBadRequest)
			return
	}
    words, err := openai.FetchAnswers(clue)
    if err != nil {
        log.Printf("Error calling OpenAI API: %v", err)
				http.Error(w, "Failed to generate answers", http.StatusInternalServerError)
        return
    }

    response := GenerateAnswersResponse{
			Answers: words,
			Clue:    clue,
		}

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}