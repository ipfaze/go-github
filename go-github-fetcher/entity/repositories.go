package entity

/*
Partial entity sent back from the GitHub API listing public repositories
(warning : some fields are missing)
*/
type RepositoriesResponse struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []Repository `json:"items"`
}

/*
Partial entity sent back from the GitHub API defining a repositories
(warning : some fields are missing)
*/
type Repository struct {
	FullName string  `json:"full_name"`
	Name     string  `json:"name"`
	Owner    Owner   `json:"owner"`
	Size     int     `json:"size"`
	Language string  `json:"language"`
	License  License `json:"license"`
}

/*
Partial entity sent back from the GitHub API defining an Owner
(warning : some fields are missing)
*/
type Owner struct {
	Login string `json:"login"`
}

type License struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Url  string `json:"url"`
}
