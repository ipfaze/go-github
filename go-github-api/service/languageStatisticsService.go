package service

import (
	Repo "go-github-api/repository"

	"encoding/json"
	"net/http"
)

// LanguageStatisticsService define all functions of the service
type LanguageStatisticsService interface {
	GetAllLanguageStatistics(w http.ResponseWriter, r *http.Request)
}

// GetAllLanguageStatistics retrieve all the language statistics by calling the repository layer
func GetAllLanguageStatistics(w http.ResponseWriter, r *http.Request) {
	statistics := Repo.GetAllLanguageStatistics()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(statistics)
}
