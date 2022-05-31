package entity

type LanguageStatistics struct {
	Language string `bson:"language"`
	Total    int    `bson:"total"`
}
