package repository

import (
	Entity "go-github-fetcher/entity"
	Helpers "go-github-fetcher/helpers"
	Mongo "go-github-fetcher/mongo"

	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type LicenseStatRepo interface {
	InsertManyLicenseStatistics(licenseStats []Entity.LicenseStatistics)
	RemoveAllLicenseStatistics()
}

//InsertManyLicenseStatistics insert many license statistics into mongodb
func InsertManyLicenseStatistics(licenseStats []Entity.LicenseStatistics) {
	RemoveAllLicenseStatistics()

	mongoClient := Mongo.Client()
	defer Mongo.Disconnect(mongoClient)

	licenseStatisticsCollection := mongoClient.Database("go-github").Collection("license_statistics")

	_, err := licenseStatisticsCollection.InsertMany(context.TODO(), Helpers.LicenseStatisticsSliceToInterfaceSlice(licenseStats))

	if err != nil {
		log.Fatal("error : inserting license statistics throw -> ", err)
	}
}

//RemoveAllLicenseStatistics remove all license statistics already existing into mongodb
func RemoveAllLicenseStatistics() {
	mongoClient := Mongo.Client()
	defer Mongo.Disconnect(mongoClient)

	licenseStatisticsCollection := mongoClient.Database("go-github").Collection("license_statistics")

	_, err := licenseStatisticsCollection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("error : deleting license statistics throw -> ", err)
	}
}
