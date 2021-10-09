package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func postsEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	if r.Method == "POST" {

		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
		client, err := mongo.Connect(context.TODO(), clientOptions)
		if err != nil {
			fmt.Println("Mongo.connect() error ", err)
			// write code to exit
		}
		// processing raw request query
		r.ParseForm()

		// setting timeout for request
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		collection := client.Database("Instagram-Backend-API").Collection("posts")
		collection2 := client.Database("Instagram-Backend-API").Collection("users")

		var result bson.M
		var result2 bson.M
		found := collection2.FindOne(context.TODO(), bson.D{{"Email", r.Form["Email"][0]}}).Decode(&result2)
		if found != nil {
			// ErrNoDocuments means that the filter did not match any documents in the collection
			if found == mongo.ErrNoDocuments {
				json.NewEncoder(w).Encode("No such user exists")
				return
			}
		}
		// ErrNoDocuments means that the filter did not match any documents in the collection
		var post Posts
		itemCount, _ := collection.CountDocuments(ctx, bson.M{})
		post.PostId = uint64(itemCount + 1)
		post.Email = r.Form["Email"][0]
		post.Caption = r.Form["Caption"][0]
		post.Image_URL = r.Form["Image_URL"][0]
		post.Posted_Timestamp = r.Form["Posted_Timestamp"][0]
		ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

		collection.InsertOne(ctx, post)

		json.NewEncoder(w).Encode(result)
		return
	} else {
		json.NewEncoder(w).Encode("Kindly make post requests on this URL to create new users.")
	}
}
