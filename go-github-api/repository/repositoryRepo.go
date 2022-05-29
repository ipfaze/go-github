package repository

import (
	Entity "go-github-api/entity"
	Mongo "go-github-api/mongo"

	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type RepositoryRepo interface {
	GetAllRepositories() []Entity.Repository
	GetByLanguage(language string) []Entity.Repository
	GetByLicense(license string) []Entity.Repository
}

//GetAllRepositories retrieve all repositories from mongodb
func GetAllRepositories() []Entity.Repository {
	mongoClient := Mongo.Client()
	defer Mongo.Disconnect(mongoClient)

	repositoriesCollection := mongoClient.Database("go-github").Collection("repositories")

	cur, err := repositoriesCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal("error : retrieving cursor of all repositories from mongodb throw -> ", err)
	}

	defer cur.Close(context.TODO())

	var repositories []Entity.Repository
	for cur.Next(context.TODO()) {
		var result Entity.Repository

		err := cur.Decode(&result)

		if err != nil {
			log.Fatal("error : parsing results of all repositories throw -> ", err)
		} else {
			repositories = append(repositories, result)
		}
	}

	return repositories
}

//GetRepositoriesByLanguage retrieve all repositories from mongodb filter by language
func GetRepositoriesByLanguage(language string) []Entity.Repository {
	mongoClient := Mongo.Client()
	defer Mongo.Disconnect(mongoClient)

	repositoriesCollection := mongoClient.Database("go-github").Collection("repositories")

	regex := bson.D{{Key: "language", Value: bson.D{
		{Key: "$regex", Value: "^" + language + ".*"},
		{Key: "$options", Value: "i"},
	}}}

	cur, err := repositoriesCollection.Find(context.TODO(), regex)
	if err != nil {
		log.Fatal("error : retrieving cursor of all repositories filtered by language throw -> ", err)
	}

	defer cur.Close(context.TODO())

	var repositories []Entity.Repository
	for cur.Next(context.TODO()) {
		var result Entity.Repository

		err := cur.Decode(&result)

		if err != nil {
			log.Fatal("error : parsing results of all repositories filtered by language throw -> ", err)
		} else {
			repositories = append(repositories, result)
		}
	}

	return repositories
}

//GetRepositoriesByLicense retrieve all repositories from mongodb filter by license
func GetRepositoriesByLicense(license string) []Entity.Repository {
	mongoClient := Mongo.Client()
	defer Mongo.Disconnect(mongoClient)

	repositoriesCollection := mongoClient.Database("go-github").Collection("repositories")

	regex := bson.D{{Key: "license.name", Value: bson.D{
		{Key: "$regex", Value: ".*" + license + ".*"},
		{Key: "$options", Value: "i"},
	}}}

	cur, err := repositoriesCollection.Find(context.TODO(), regex)
	if err != nil {
		log.Fatal("error : retrieving cursor of all repositories filtered by license throw -> ", err)
	}

	defer cur.Close(context.TODO())

	var repositories []Entity.Repository
	for cur.Next(context.TODO()) {
		var result Entity.Repository

		err := cur.Decode(&result)

		if err != nil {
			log.Fatal("error : parsing results of all repositories filtered by license throw -> ", err)
		} else {
			repositories = append(repositories, result)
		}
	}

	return repositories
}
