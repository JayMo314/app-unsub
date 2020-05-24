package mongoDb

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserLink struct {
	UserId    int    `json:"user_id"`
	UnsubLink string `json:"unsub_link"`
}

func UserLinkAdd(ul UserLink) {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb+srv://app-user:app-pass@email-unsubscribe-b5iub.mongodb.net/test?retryWrites=true&w=majority")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		fmt.Println(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	col := client.Database("email-unsubscribe").Collection("links")
	fmt.Println("Collection type:", reflect.TypeOf(col))

	result, insertErr := col.InsertOne(ctx, ul)

	if insertErr != nil {
		fmt.Println(insertErr)
	} else {
		fmt.Println("UserId", result.InsertedID)
	}
}
