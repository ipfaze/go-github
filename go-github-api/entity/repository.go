package entity

type Repository struct {
	FullName   string  `json:"full_name" bson:"full_name"`
	Owner      string  `json:"owner" bson:"owner"`
	Repository string  `json:"repository" bson:"repository"`
	Language   string  `json:"language" bson:"language"`
	Bytes      int     `json:"bytes" bson:"bytes"`
	License    License `json:"license" bson:"license"`
}

type License struct {
	Key  string `json:"key" bson:"key"`
	Name string `json:"name" bson:"name"`
	Url  string `json:"url" bson:"url"`
}
