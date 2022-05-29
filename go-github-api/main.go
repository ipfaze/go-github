package main

import (
	Mongo "go-github-api/mongo"
	Repo "go-github-api/repository"

	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func main() {
	// Check mongodb connection
	Mongo.Ping()

	// Endpoint for repositories
	http.HandleFunc("/api/repositories", getRepos)
	http.HandleFunc("/api/repositories/language/", getReposByLanguage)
	http.HandleFunc("/api/repositories/license/", getReposByLicense)

	// Endpoint for statistics
	http.HandleFunc("/api/statistics/language", getLanguageStatistics)
	http.HandleFunc("/api/statistics/license", getLicenseStatistics)

	log.Println("info : server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getRepos(w http.ResponseWriter, r *http.Request) {
	repositories := Repo.GetAllRepositories()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories)
}

func getReposByLanguage(w http.ResponseWriter, r *http.Request) {
	language := strings.TrimPrefix(r.URL.Path, "/api/repositories/language/")

	repositories := Repo.GetRepositoriesByLanguage(language)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories)
}

func getReposByLicense(w http.ResponseWriter, r *http.Request) {
	license := strings.TrimPrefix(r.URL.Path, "/api/repositories/license/")

	repositories := Repo.GetRepositoriesByLicense(license)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories)
}

func getLanguageStatistics(w http.ResponseWriter, r *http.Request) {
	statistics := Repo.GetAllLanguageStatistics()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(statistics)
}

func getLicenseStatistics(w http.ResponseWriter, r *http.Request) {
	statistics := Repo.GetAllLicenseStatistics()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(statistics)
}
