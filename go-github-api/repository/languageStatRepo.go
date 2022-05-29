package repository

import (
	Entity "go-github-api/entity"
	Mongo "go-github-api/mongo"

	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type LanguageStatisticsRepo interface {
	GetAllLanguageStatistics() []Entity.LanguageStatistics
}

//GetAllLanguageStatistics retrieve a bunch of language statistics from repositories saved into mongodb
func GetAllLanguageStatistics() []Entity.LanguageStatistics {
	mongoClient := Mongo.Client()
	defer Mongo.Disconnect(mongoClient)

	languageStatisticsCollection := mongoClient.Database("go-github").Collection("language_statistics")

	cur, err := languageStatisticsCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal("error : retrieving cursor of all language statistics from mongodb throw -> ", err)
	}

	defer cur.Close(context.TODO())

	var languageStatistics []Entity.LanguageStatistics
	for cur.Next(context.TODO()) {
		var result Entity.LanguageStatistics

		err := cur.Decode(&result)

		if err != nil {
			log.Fatal("error : parsing results of all language statistics throw -> ", err)
		} else {
			languageStatistics = append(languageStatistics, result)
		}
	}

	return languageStatistics
}
