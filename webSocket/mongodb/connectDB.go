package mongodb



import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectDB() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
}

func GetClient() *mongo.Client {
	if client == nil {
		ConnectDB()
	}
	return client
}

func GetCollection(name string) *mongo.Collection {
	return GetClient().Database("ChatAppInGo").Collection(name)
}