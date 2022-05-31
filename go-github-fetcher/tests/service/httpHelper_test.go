package helper_test

import (
	Service "go-github-fetcher/service"
	"reflect"
	"time"

	"testing"
)

func TestHttpClient(t *testing.T) {
	httpClient := Service.HttpClient()

	if reflect.TypeOf(httpClient).String() != "*http.Client" {
		t.Error("The type return should be http.Client but get : ", reflect.TypeOf(httpClient))
	}
}

func TestHttpCallInvalidUrl(t *testing.T) {
	url := "myinvalidurl"

	_, err := Service.HttpCall(url)

	if err == nil {
		t.Error("An invalid url should throw an error")
	}
}

func TestHttpCallValidUrlButNotExist(t *testing.T) {
	url := "https://api.github.com/my/valid/url"

	_, err := Service.HttpCall(url)

	if err == nil {
		t.Error("The url : '", url, "' is valid but dos not exist, this should throw an error")
	}
}

func TestHttpCallValidUrlAndExist(t *testing.T) {
	url := "https://github.com/"

	response, err := Service.HttpCall(url)

	if err != nil {
		t.Error("This function should not throw an error but get : ", err)
	}
	if len(response) == 0 {
		t.Error("The response should not be empty")
	}
}

func TestConstructGitHubUrl(t *testing.T) {
	githubUrl := Service.ConstructGitHubUrl()

	date := time.Now().Add(-24 * time.Hour)
	expectedGithubUrl := "https://api.github.com/search/repositories?q=created%3A>" + date.Format("2006-01-02") + "&sort=updated&order=desc&per_page=100"

	if githubUrl != expectedGithubUrl {
		t.Error("The Github url was not like expected.\nExpected : ", expectedGithubUrl, "\nGet : ", githubUrl)
	}
}
