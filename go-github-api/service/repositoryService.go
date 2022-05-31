package service

import (
	Repo "go-github-api/repository"

	"encoding/json"
	"net/http"
	"strings"
)

// RepositoryService  define all functions of the service
type RepositoryService interface {
	GetAllRepositories(w http.ResponseWriter, r *http.Request)
	GetRepositoriesByLanguage(w http.ResponseWriter, r *http.Request)
	GetRepositoriesByLicense(w http.ResponseWriter, r *http.Request)
}

// GetAllRepositories retrieve all the repositories by calling the repository layer
func GetAllRepositories(w http.ResponseWriter, r *http.Request) {
	repositories := Repo.GetAllRepositories()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories)
}

// GetRepositoriesByLanguage retrieve all the repositories filtered by language define in url by calling the repository layer
func GetRepositoriesByLanguage(w http.ResponseWriter, r *http.Request) {
	language := strings.TrimPrefix(r.URL.Path, "/api/repositories/language/")

	repositories := Repo.GetRepositoriesByLanguage(language)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories)
}

// GetRepositoriesByLicense retrieve all the repositories filtered by license define in url by calling the repository layer
func GetRepositoriesByLicense(w http.ResponseWriter, r *http.Request) {
	license := strings.TrimPrefix(r.URL.Path, "/api/repositories/license/")

	repositories := Repo.GetRepositoriesByLicense(license)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories)
}
