package helpers

import (
	"io/ioutil"
	"log"
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
func HttpCall(url string) []byte {
	client := HttpClient()

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

//ConstructGitHubUrl construct the GitHub API url for repositories
func ConstructGitHubUrl() string {
	date := time.Now().Add(-24 * time.Hour)

	q := "created%3A>" + date.Format("2006-01-02")
	sort := "updated"
	order := "desc"
	per_page := "100"

	return GITHUB_API_URL + "?q=" + q + "&sort=" + sort + "&order=" + order + "&per_page=" + per_page
}
