package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func usersEndpoint(w http.ResponseWriter, r *http.Request) {
	// SETTING HEADERS ON RESPONSE
	w.Header().Add("content-type", "application/json")

	// ESTABLISHING CONNECTION WITH DATABASE
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() error ", err)
	}

	// MAKES CONCURRENCY A TAD BETTER
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	collection := client.Database("Instagram-Backend-API").Collection("users")

	// HANDLING POST REQUESTS
	if r.Method == "POST" {
		// PROCESSING RAW QUERY
		r.ParseForm()
		var result bson.M
		found := collection.FindOne(context.TODO(), bson.D{{"Email", r.Form["Email"][0]}}).Decode(&result)
		if found != nil {
			// ErrNoDocuments means that the filter did not match any documents in the collection
			if found == mongo.ErrNoDocuments {
				var user User
				itemCount, _ := collection.CountDocuments(ctx, bson.M{})
				user.UserId = uint64(itemCount + 1)
				user.Name = r.Form["Name"][0]
				user.Email = r.Form["Email"][0]
				myslice := make([]uint64, 0)
				user.PostSlice = myslice
				// ENCRYPTING AND STORING PASSWORD
				user.Password = givePwdHash(r.Form["Password"][0])
				ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)

				// sends identification key as response
				result, _ := collection.InsertOne(ctx, user)
				json.NewEncoder(w).Encode(result)
				return
			}
			log.Fatal(found)
		}
		json.NewEncoder(w).Encode("User already exists.")
		return
	} else {
		uids, ok := r.URL.Query()["uid"]

		if !ok || len(uids[0]) < 1 {
			log.Println("Url Param 'uid' is missing")
			return
		}
		uid, _ := strconv.ParseInt(uids[0], 0, 64)
		JSONData := struct {
			UserId uint64 `bson:"User_id"`
			Name   string `bson:"Name"`
			Email  string `bson:"Email"`
		}{}
		collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "User_id", Value: uid}}).Decode(&JSONData)
		fmt.Println(JSONData)
		json.NewEncoder(w).Encode(JSONData)
		return
	}
}
