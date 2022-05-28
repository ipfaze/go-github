package entity

type LanguageStatistics struct {
	Language string `json:"language"`
	Total    int    `json:"total"`
}

type LicenseStatistics struct {
	License string `json:"license"`
	Total   int    `json:"total"`
}
