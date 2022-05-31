package repository

import (
	Entity "go-github-fetcher/entity"
	Mongo "go-github-fetcher/mongo"

	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RepositoryRepo interface {
	InsertRepository(repository Entity.Repository) string
	RemoveAllRepositoryBefore(objectId string)
}

//InsertRepository insert a new Repository into mongodb
func InsertRepository(repository Entity.Repository) string {
	mongoClient := Mongo.Client()
	defer Mongo.Disconnect(mongoClient)

	repositoriesCollection := mongoClient.Database("go-github").Collection("repositories")

	bsonObject := bson.D{
		{Key: "full_name", Value: repository.FullName},
		{Key: "owner", Value: repository.Owner.Login},
		{Key: "repository", Value: repository.Name},
		{Key: "language", Value: repository.Language},
		{Key: "bytes", Value: repository.Size},
		{Key: "license", Value: repository.License},
	}

	res, err := repositoriesCollection.InsertOne(context.TODO(), bsonObject)

	if err != nil {
		log.Fatal("error : inserting repository throw -> ", err)
	}

	return res.InsertedID.(primitive.ObjectID).Hex()
}

//RemoveAllRepository remove all Repository into mongodb before the objectId set in parameter
func RemoveAllRepositoryBefore(objectId string) {
	mongoClient := Mongo.Client()
	defer Mongo.Disconnect(mongoClient)

	oid, err := primitive.ObjectIDFromHex(objectId)
	if err != nil {
		log.Fatal("error : convert primitive.ObjectID throw -> ", err)
	}

	repositoriesCollection := mongoClient.Database("go-github").Collection("repositories")

	_, err = repositoriesCollection.DeleteMany(context.TODO(), bson.M{"_id": bson.M{"$lt": oid}})
	if err != nil {
		log.Fatal("error : deleting repository before ", oid, " throw -> ", err)
	}
}
