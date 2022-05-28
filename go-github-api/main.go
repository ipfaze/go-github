package main

import (
	Mongo "go-github-api/mongo"
	"strings"

	"encoding/json"
	"log"
	"net/http"
)

func main() {
	Mongo.Ping()

	http.HandleFunc("/repos", getRepos)
	http.HandleFunc("/repos/language/", getReposByLanguage)
	http.HandleFunc("/repos/license/", getReposByLicense)

	log.Println("info : server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getRepos(w http.ResponseWriter, r *http.Request) {
	repositories := Mongo.GetRepos()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories)
}

func getReposByLanguage(w http.ResponseWriter, r *http.Request) {
	language := strings.TrimPrefix(r.URL.Path, "/repos/language/")

	repositories := Mongo.GetReposByLanguage(language)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories)
}

func getReposByLicense(w http.ResponseWriter, r *http.Request) {
	license := strings.TrimPrefix(r.URL.Path, "/repos/license/")

	repositories := Mongo.GetReposByLicense(license)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories)
}