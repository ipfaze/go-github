package mongo

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	MONGO_URL = "mongodb://" + os.Getenv("MONGO_USER") + ":" + os.Getenv("MONGO_PASSWORD") + "@" + os.Getenv("MONGO_HOST") + ":" + os.Getenv("MONGO_PORT") + "/?maxPoolSize=20&w=majority"
)

//Client create a mongodb connection and return it
func Client() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGO_URL))
	if err != nil {
		log.Fatal("error : creating mongodb client throw -> ", err)
	}

	return client
}

//Disconnect release the mongodb connection
func Disconnect(client *mongo.Client) {
	err := client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal("error : disconnecting mongodb client throw -> ", err)
	}
}

//Ping test the connection to mongodb
func Ping() {
	mongoClient := Client()
	defer Disconnect(mongoClient)

	err := mongoClient.Ping(context.TODO(), readpref.Primary())

	if err != nil {
		log.Fatal("error : ping database throw -> ", err)
	}

	log.Println("info : successfully connected and pinged")
}
