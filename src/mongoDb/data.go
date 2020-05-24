package mongoDb

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type Trainer struct {
	Name string
	Age  int
	City string
}

type Document struct {
	userId    int    `json:"user_id"`
	unsubLink string `json:"unsub_link"`
}

func Access(document Document) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://app-user:app-pass@email-unsubscribe-b5iub.mongodb.net/test?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	ctx, _

	if err != nil {
		log.Fatal(err)
	}

	col := client.Database("email-unsubscribe").Collection("links")
	fmt.Println("Collection type:", reflect.TypeOf(col))

	col.InsertOne()

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
}
