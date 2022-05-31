package service

import (
	Repo "go-github-api/repository"

	"encoding/json"
	"net/http"
)

// LicenseStatisticsService define all functions of the service
type LicenseStatisticsService interface {
	GetAllLicenseStatistics(w http.ResponseWriter, r *http.Request)
}

// GetAllLicenseStatistics retrieve all the license statistics by calling the repository layer
func GetAllLicenseStatistics(w http.ResponseWriter, r *http.Request) {
	statistics := Repo.GetAllLicenseStatistics()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(statistics)
}
