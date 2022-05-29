package entity

type LicenseStatistics struct {
	License string `bson:"license"`
	Total   int    `bson:"total"`
}
