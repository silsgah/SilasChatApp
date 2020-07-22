package main

import (
	"fmt"

	"github.com/gin-contrib/sessions/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var uri = "mongodb://localhost:27017"
	var client *mongo.Client
	var err error
	client, err = mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		fmt.Println("error: ", err)
	}

	// clientOptions := options.Client().ApplyURI("mongo://mongodb:27017")
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// // client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = client.Ping(context.TODO(), nil)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("connected to mongoDB")
	// ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer client.Disconnect(ctx)
}
