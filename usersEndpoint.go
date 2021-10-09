package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func usersEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() error ", err)
		// write code to exit
	}
	// setting timeout for request
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("Instagram-Backend-API").Collection("users")
	if r.Method == "POST" {
		// processing raw request query
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
	} else {
		uids, ok := r.URL.Query()["uid"]

		if !ok || len(uids[0]) < 1 {
			log.Println("Url Param 'uid' is missing")
			return
		}
		uid, _ := strconv.ParseInt(uids[0], 0, 64)
		fmt.Println(reflect.TypeOf(uid))
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
