package repository

import (
	Entity "go-github-fetcher/entity"
	Helpers "go-github-fetcher/helpers"
	Mongo "go-github-fetcher/mongo"

	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type LanguageStatRepo interface {
	InsertManyLanguageStatistics(languageStats []Entity.LanguageStatistics)
	RemoveAllLanguageStatistics()
}

//InsertManyLanguageStatistics insert many language statistics into mongodb
func InsertManyLanguageStatistics(languageStats []Entity.LanguageStatistics) {
	mongoClient := Mongo.Client()
	defer Mongo.Disconnect(mongoClient)

	languageStatisticsCollection := mongoClient.Database("go-github").Collection("language_statistics")

	_, err := languageStatisticsCollection.InsertMany(context.TODO(), Helpers.LanguageStatisticsSliceToInterfaceSlice(languageStats))

	if err != nil {
		log.Fatal("error : inserting language statistics throw -> ", err)
	}
}

//RemoveAllLanguageStatistics remove all language statistics already existing into mongodb
func RemoveAllLanguageStatistics() {
	mongoClient := Mongo.Client()
	defer Mongo.Disconnect(mongoClient)

	languageStatisticsCollection := mongoClient.Database("go-github").Collection("language_statistics")

	_, err := languageStatisticsCollection.DeleteMany(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("error : deleting language statistics throw -> ", err)
	}
}
