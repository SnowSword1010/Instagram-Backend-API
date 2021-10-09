package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func usersEndpoint(w http.ResponseWriter, r *http.Request) {
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
		w.Header().Add("content-type", "application/json")
		collection := client.Database("Instagram-Backend-API").Collection("users")
		var result bson.M
		found := collection.FindOne(context.TODO(), bson.D{{"Email", r.Form["Email"][0]}}).Decode(&result)
		if found != nil {
			// ErrNoDocuments means that the filter did not match any documents in the collection
			if found == mongo.ErrNoDocuments {
				var user User
				itemCount, _ := collection.CountDocuments(ctx, bson.M{})
				user.UserId = uint64(itemCount + 1)
				fmt.Println(user.UserId)
				user.Name = r.Form["Name"][0]
				user.Email = r.Form["Email"][0]
				// ENCRYPTING AND STORING PASSWORD
				user.Password = givePwdHash(r.Form["Password"][0])
				myslice := make([]uint64, 0)
				user.PostID = myslice
				ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

				//this code sends identification key as response
				result, _ := collection.InsertOne(ctx, user)
				json.NewEncoder(w).Encode(result)
				return
			}
			log.Fatal(found)
		}
		json.NewEncoder(w).Encode("User already exists.")
		return
	} else if r.Method == "GET" {
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode("Kindly make post requests on this URL to create new users.")
	}
}
