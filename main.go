package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	_id      primitive.ObjectID `json:"_id" bson:"_id"`
	UserId   uint64             `json:"user_id" bson:"user_id"`
	Name     string             `json:"Name" bson:"Name"`
	Email    string             `json:"Email" bson:"Email"`
	Password string             `json:"Password" bson:"Password"`
	PostID   []uint64           `json:"PostID" bson:"PostID"`
}

type Posts struct {
	id               primitive.ObjectID `json:"_id" bson:"_id"`
	Caption          string             `json:"Caption" bson:"Caption"`
	Image_URL        string             `json:"Image_URL" bson:"Image_URL"`
	Posted_Timestamp string             `json:"Posted_Timestamp" bson:"Posted_Timestamp"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode("Welcome to Instagram Backend API | Appointy Task")
}

func main() {
	fmt.Println("Starting the server at port 8080")

	// client, _ = mongo.Connect(ctx, "mongodb://localhost:27017")
	http.HandleFunc("/users", usersEndpoint)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
