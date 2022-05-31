package main

import (
	Mongo "go-github-api/mongo"
	Service "go-github-api/service"

	"log"
	"net/http"
)

func main() {
	// Check mongodb connection
	Mongo.Ping()

	// Endpoint for repositories
	http.HandleFunc("/api/repositories", Service.GetAllRepositories)
	http.HandleFunc("/api/repositories/language/", Service.GetRepositoriesByLanguage)
	http.HandleFunc("/api/repositories/license/", Service.GetRepositoriesByLicense)

	// Endpoint for statistics
	http.HandleFunc("/api/statistics/language", Service.GetAllLanguageStatistics)
	http.HandleFunc("/api/statistics/license", Service.GetAllLicenseStatistics)

	log.Println("info : server listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
