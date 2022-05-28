package mongo

import (
	Entity "go-github-fetcher/entity"

	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MONGO_URL = "mongodb://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT") + "/?maxPoolSize=20&w=majority"
)

/*
Function to create a mongodb connection and return it
*/
func Client() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_URL))
	if err != nil {
		log.Fatal("error : creating mongodb client throw -> ", err)
	}

	return client
}

/*
Function to release the mongodb connection
*/
func Disconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal("error : disconnecting mongodb client throw -> ", err)
	}
}

/*
Function to test the connection to mongodb
*/
func Ping() {
	mongoClient := Client()
	defer Disconnect(mongoClient)

	err := mongoClient.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		log.Fatal("error : ping database throw -> ", err)
	}

	log.Println("info : successfully connected and pinged")
}

/*
Function to insert a new GitItem into mongodb
*/
func InsertRepository(repository Entity.Repository) string {
	mongoClient := Client()
	defer Disconnect(mongoClient)

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

/*
Function to remove all GitItem into mongodb before the objectId set in parameter
*/
func RemoveAllRepository(objectId string) {
	mongoClient := Client()
	defer Disconnect(mongoClient)

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