package entity

//RepositoriesResponse (partial) entity sent back from the GitHub API listing public repositories
//(warning : some fields are missing)
type RepositoriesResponse struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []Repository `json:"items"`
}

//Repository (partial) entity sent back from the GitHub API
//(warning : some fields are missing)
type Repository struct {
	FullName string  `json:"full_name"`
	Name     string  `json:"name"`
	Owner    Owner   `json:"owner"`
	Size     int     `json:"size"`
	Language string  `json:"language"`
	License  License `json:"license"`
}

//Owner (partial) entity sent back from the GitHub API
//(warning : some fields are missing)
type Owner struct {
	Login string `json:"login"`
}

//License sent back from the GitHub API
type License struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
