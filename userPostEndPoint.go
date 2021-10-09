package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func userPostEndPoint(w http.ResponseWriter, r *http.Request) {
	// SETTING HEADERS ON RESPONSE
	w.Header().Add("content-type", "application/json")

	//ESTABLISHING CONNECTION WITH DATABASE
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() error ", err)
	}

	// DEFINING COLLECTIONS ; collection - posts database ; collection2 - users database
	collection := client.Database("Instagram-Backend-API").Collection("posts")
	collection2 := client.Database("Instagram-Backend-API").Collection("users")

	// HANDLING GET REQUESTS
	if r.Method == "GET" {
		uids, ok := r.URL.Query()["uid"]

		if !ok || len(uids[0]) < 1 {
			log.Println("Url Param 'uid' is missing")
			return
		}

		uid, _ := strconv.ParseInt(uids[0], 0, 64)

		// implementing pagination : showing 5 posts at a time
		pages, notok := r.URL.Query()["page"]

		if !notok || len(pages[0]) < 1 {
			log.Println("Page Param is missing")
			return
		}

		page, _ := strconv.ParseInt(pages[0], 0, 64)

		// ALGORITHM TO IMPLEMENT PAGINATION
		var startindx, endindx int64 = page * 5, page*5 + 5

		if startindx <= -1 || endindx <= -1 {
			return
		}

		JSONData := struct {
			UserId    uint64   `bson:"User_id"`
			Name      string   `bson:"Name"`
			Email     string   `bson:"Email"`
			PostSlice []uint64 `json:"PostSlice" bson:"PostSlice"`
		}{}

		collection2.FindOne(context.TODO(), bson.D{primitive.E{Key: "User_id", Value: uid}}).Decode(&JSONData)

		mySlice := JSONData.PostSlice
		JSONPost := struct {
			Post_id          uint64 `bson:"Post_id"`
			Caption          string `bson:"Caption"`
			Image_URL        string `bson:"Image_URL"`
			Posted_Timestamp string `bson:"Posted_Timestamp"`
		}{}

		// VITAL CHECK TO PREVENT SOCKET CRASH
		if endindx > int64(cap(mySlice)) {
			endindx = int64(cap(mySlice))
		}

		// DISPLAYING JSON RESPONSES OF RELEVANT POSTS
		for index := startindx; index < endindx; index++ {
			element := mySlice[index]
			fmt.Println(element)
			fmt.Println(element)
			collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "Post_id", Value: element}}).Decode(&JSONPost)
			json.NewEncoder(w).Encode(JSONPost)
		}
	} else {
		// CAUTION MESSAGE FOR WRONG REQUEST TYPES
		json.NewEncoder(w).Encode("Kindly use only GET requests on this route")
	}
}
