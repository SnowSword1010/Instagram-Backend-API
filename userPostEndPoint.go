package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func userPostEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() error ", err)
		// write code to exit
	}
	// setting timeout for request
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	collection := client.Database("Instagram-Backend-API").Collection("posts")
	if r.Method == "GET" {
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
		// json.NewEncoder(w).Encode("Kindly make post requests on this URL to create new users.")
	}
}
