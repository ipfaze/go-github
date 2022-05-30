package helpers

import (
	"encoding/json"

	Entity "go-github-fetcher/entity"
)

//ConvertBytesToRepositoriesResponse convert a byte array to the RepositoriesResponse struct
func ConvertBytesToRepositoriesResponse(resp []byte) Entity.RepositoriesResponse {
	var gitResponse Entity.RepositoriesResponse

	err := json.Unmarshal(resp, &gitResponse)
	if err != nil {
		return Entity.RepositoriesResponse{}
	}

	return gitResponse
}
