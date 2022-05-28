package entity

type LanguageStatistics struct {
	Language string `bson:"language"`
	Total    int    `bson:"total"`
}

type LicenseStatistics struct {
	License string `bson:"license"`
	Total   int    `bson:"total"`
}
