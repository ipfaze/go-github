package main

import (
	Entity "go-github-fetcher/entity"
	Mongo "go-github-fetcher/mongo"

	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	GITHUB_API_URL = "https://api.github.com/search/repositories"
)

/*
Function instanciate an http client and return it
*/
func httpClient() *http.Client {
	return &http.Client{Timeout: 10 * time.Second}
}

/*
Function to prepare and send the request to the GitHub API to retrieve data about public repositories
It returns the body content of the response as an array of byte
*/
func call(url string) []byte {
	client := httpClient()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("error : creating the request throw -> ", err.Error())
		return nil
	}

	// Prepare Header to request GitHub API
	req.Header.Set("user-agent", "go-github-fetcher")
	req.Header.Add("Accept", "application/vnd.github.v3+json")

	// Execute request
	response, err := client.Do(req)
	if err != nil {
		log.Fatal("error : requesting API throw -> ", err.Error())
		return nil
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatal("error : response status code is not OK, get -> ", response.StatusCode)
		return nil
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("error : reading response body throw -> ", err.Error())
		return nil
	}

	return body
}

/*
Function to convert a byte array to the GitResponse struct
*/
func convertBytesToGitResponse(resp []byte) Entity.RepositoriesResponse {
	var gitResponse Entity.RepositoriesResponse

	err := json.Unmarshal(resp, &gitResponse)
	if err != nil {
		log.Fatal("error : unmarshaling to Entity.RepositoriesResponse throw -> ", err)
		return Entity.RepositoriesResponse{}
	}

	return gitResponse
}

/*
Function to send the request, retrieve data from the response, save the response's data to the mongodb and delete all documents before the first inserted
*/
func sendRequestAndSaveData(url string) {
	log.Println("Refreshing database ", time.Now().Format("2006-01-02 15:04:05"))

	resp := call(url)

	gitResp := convertBytesToGitResponse(resp)

	var idToDelete string

	for i := 0; i < len(gitResp.Items); i++ {
		id := Mongo.InsertRepository(gitResp.Items[i])
		if i == 0 {
			idToDelete = id
		}
	}

	Mongo.RemoveAllRepository(idToDelete)
}

/*
Function to construct the GitHub API url for repositories
*/
func constructGitHubUrl() string {
	date := time.Now().Add(-24 * time.Hour)

	q := "created%3A>" + date.Format("2006-01-02")
	sort := "updated"
	order := "desc"
	per_page := "100"

	return GITHUB_API_URL + "?q=" + q + "&sort=" + sort + "&order=" + order + "&per_page=" + per_page
}

/*
Main function executing the main logic every 6 minutes
*/
func main() {
	// Verify database is accessible
	Mongo.Ping()

	url := constructGitHubUrl()
	sendRequestAndSaveData(url)

	ticker := time.NewTicker(6 * time.Minute)
	quit := make(chan struct{})
	for {
		select {
		case <-ticker.C:
			url := constructGitHubUrl()
			sendRequestAndSaveData(url)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
