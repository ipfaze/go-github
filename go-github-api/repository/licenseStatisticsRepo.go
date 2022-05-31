package repository

import (
	Entity "go-github-api/entity"
	Mongo "go-github-api/mongo"

	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type LicenseStatisticsRepo interface {
	GetAllLicenseStatistics() []Entity.LicenseStatistics
}

//GetLicenseStatistics retrieve a bunch of license statistics from repositories saved into mongodb
func GetAllLicenseStatistics() []Entity.LicenseStatistics {
	mongoClient := Mongo.Client()
	defer Mongo.Disconnect(mongoClient)

	licenseStatisticsCollection := mongoClient.Database("go-github").Collection("license_statistics")

	cur, err := licenseStatisticsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal("error : retrieving cursor of all license statistics from mongodb throw -> ", err)
	}

	defer cur.Close(context.TODO())

	var licenseStatistics []Entity.LicenseStatistics
	for cur.Next(context.TODO()) {
		var result Entity.LicenseStatistics

		err := cur.Decode(&result)

		if err != nil {
			log.Fatal("error : parsing results of all license statistics throw -> ", err)
		} else {
			licenseStatistics = append(licenseStatistics, result)
		}
	}

	return licenseStatistics
}
