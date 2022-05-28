package mongo

import (
	Entity "go-github-api/entity"
	"log"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MONGO_URL = "mongodb://root:example@localhost:27017/?maxPoolSize=20&w=majority"
	//MONGO_URL = "mongodb://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT") + "/?maxPoolSize=20&w=majority"
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
Function to retrieve all repositories from mongodb
*/
func GetRepos() Entity.Repositories {
	mongoClient := Client()
	defer Disconnect(mongoClient)

	repositoriesCollection := mongoClient.Database("go-github").Collection("repositories")

	cur, err := repositoriesCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal("error : retrieving cursor of all repositories from mongodb throw -> ", err)
	}

	defer cur.Close(context.TODO())

	var repositories Entity.Repositories
	for cur.Next(context.TODO()) {
		var result Entity.Repository

		err := cur.Decode(&result)

		if err != nil {
			log.Fatal("error : parsing results of all repositories throw -> ", err)
		} else {
			repositories.Repositories = append(repositories.Repositories, result)
		}
	}

	return repositories
}

/*
Function to retrieve all repositories from mongodb filter by language
*/
func GetReposByLanguage(language string) Entity.Repositories {
	mongoClient := Client()
	defer Disconnect(mongoClient)

	repositoriesCollection := mongoClient.Database("go-github").Collection("repositories")

	regex := bson.D{{Key: "language", Value: bson.D{
		{Key: "$regex", Value: language},
		{Key: "$options", Value: "i"},
	}}}

	cur, err := repositoriesCollection.Find(context.TODO(), regex)
	if err != nil {
		log.Fatal("error : retrieving cursor of all repositories filtered by language throw -> ", err)
	}

	defer cur.Close(context.TODO())

	var repositories Entity.Repositories
	for cur.Next(context.TODO()) {
		var result Entity.Repository

		err := cur.Decode(&result)

		if err != nil {
			log.Fatal("error : parsing results of all repositories filtered by language throw -> ", err)
		} else {
			repositories.Repositories = append(repositories.Repositories, result)
		}
	}

	if repositories.Repositories == nil {
		repositories.Repositories = []Entity.Repository{}
	}

	return repositories
}

/*
Function to retrieve all repositories from mongodb filter by license
*/
func GetReposByLicense(license string) Entity.Repositories {
	mongoClient := Client()
	defer Disconnect(mongoClient)

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

	var repositories Entity.Repositories
	for cur.Next(context.TODO()) {
		var result Entity.Repository

		err := cur.Decode(&result)

		if err != nil {
			log.Fatal("error : parsing results of all repositories filtered by license throw -> ", err)
		} else {
			repositories.Repositories = append(repositories.Repositories, result)
		}
	}

	if repositories.Repositories == nil {
		repositories.Repositories = []Entity.Repository{}
	}

	return repositories
}
