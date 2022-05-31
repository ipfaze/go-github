package main

import (
	Service "go-github-fetcher/service"

	"time"
)

//main function executing the main logic every 6 minutes
func main() {
	url := Service.ConstructGitHubUrl()
	Service.RetrieveRepositoriesAndSaveIt(url)

	ticker := time.NewTicker(6 * time.Minute)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			url := Service.ConstructGitHubUrl()
			Service.RetrieveRepositoriesAndSaveIt(url)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
