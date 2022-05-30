package helpers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	GITHUB_API_URL = "https://api.github.com/search/repositories"
)

//HttpClient instanciate an http client and return it
func HttpClient() *http.Client {
	return &http.Client{Timeout: 10 * time.Second}
}

//HttpCall prepare and send the request to the GitHub API to retrieve data about public repositories
//It returns the body content of the response as an array of byte
func HttpCall(url string) ([]byte, error) {
	client := HttpClient()

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Prepare Header to request GitHub API
	req.Header.Set("user-agent", "go-github-fetcher")
	req.Header.Add("Accept", "application/vnd.github.v3+json")

	// Execute request
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	fmt.Println(response.StatusCode != http.StatusOK)
	if response.StatusCode != http.StatusOK {
		errorMessage := "The response status code for " + url + " is not 200"
		return nil, errors.New(errorMessage)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

//ConstructGitHubUrl construct the GitHub API url for repositories
func ConstructGitHubUrl() string {
	date := time.Now().Add(-24 * time.Hour)

	q := "created%3A>" + date.Format("2006-01-02")
	sort := "updated"
	order := "desc"
	per_page := "100"

	return GITHUB_API_URL + "?q=" + q + "&sort=" + sort + "&order=" + order + "&per_page=" + per_page
}
