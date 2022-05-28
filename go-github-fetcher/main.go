package main

import (
	Entity "go-github-fetcher/entity"
	Helpers "go-github-fetcher/helpers"
	Mongo "go-github-fetcher/mongo"

	"log"
	"time"
)

//main function executing the main logic every 6 minutes
func main() {
	// Verify database is accessible
	Mongo.Ping()

	url := Helpers.ConstructGitHubUrl()
	sendRequestAndSaveData(url)

	ticker := time.NewTicker(6 * time.Minute)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			url := Helpers.ConstructGitHubUrl()
			sendRequestAndSaveData(url)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

//sendRequestAndSaveData send the request, retrieve data from the response, save the response's data to the mongodb and delete all documents before the first inserted
func sendRequestAndSaveData(url string) {
	log.Println("Refreshing database ", time.Now().Format("2006-01-02 15:04:05"))

	resp := Helpers.HttpCall(url)

	repositoriesResponse := Helpers.ConvertBytesToRepositoriesResponse(resp)

	var idToDelete string
	var languageStats []Entity.LanguageStatistics
	var licenseStats []Entity.LicenseStatistics
	var key int

	for i := 0; i < len(repositoriesResponse.Items); i++ {
		key = Helpers.FindLanguageStatistics(languageStats, repositoriesResponse.Items[i].Language)
		if key == -1 {
			languageStats = append(languageStats, Entity.LanguageStatistics{Language: repositoriesResponse.Items[i].Language, Total: 1})
		} else {
			languageStats[key].Total++
		}

		key = Helpers.FindLicenseStatistics(licenseStats, repositoriesResponse.Items[i].License.Name)
		if key == -1 {
			licenseStats = append(licenseStats, Entity.LicenseStatistics{License: repositoriesResponse.Items[i].License.Name, Total: 1})
		} else {
			licenseStats[key].Total++
		}

		id := Mongo.InsertRepository(repositoriesResponse.Items[i])
		if i == 0 {
			idToDelete = id
		}
	}

	Mongo.RemoveAllRepository(idToDelete)

	Mongo.InsertLanguageStatistics(languageStats)
	Mongo.InsertLicenseStatistics(licenseStats)
}
