package helper_test

import (
	Entity "go-github-fetcher/entity"
	Helper "go-github-fetcher/helpers"

	"bytes"
	"encoding/json"
	"reflect"
	"testing"
)

func TestShouldConvertBytesToRepositoriesResponse(t *testing.T) {
	myRepositoriesResponse := Entity.RepositoriesResponse{TotalCount: 1, IncompleteResults: false, Items: []Entity.Repository{}}

	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(myRepositoriesResponse)
	repositoriesResponseBytes := reqBodyBytes.Bytes()

	repositoryResponseConverted := Helper.ConvertBytesToRepositoriesResponse(repositoriesResponseBytes)
	if !reflect.DeepEqual(repositoryResponseConverted, myRepositoriesResponse) {
		t.Error("The RepositoriesResponse should be the same than before convertion")
	}
}

func TestShouldNotConvertBytesToRepositoriesResponse(t *testing.T) {
	emptyRepositoriesResponse := Entity.RepositoriesResponse{}
	emptyBytes := []byte{}

	myRepositoryResponse := Helper.ConvertBytesToRepositoriesResponse(emptyBytes)
	if !reflect.DeepEqual(myRepositoryResponse, emptyRepositoriesResponse) {
		t.Error("An empty byte slice should return an empty RepositoriesResponse")
	}
}
