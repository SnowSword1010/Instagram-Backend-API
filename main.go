package main

import (
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	id       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"Name" bson:"Name"`
	Email    string             `json:"Email" bson:"Email"`
	Password string             `json:"Password" bson:"Password"`
}

type Posts struct {
	id               primitive.ObjectID `json:"_id" bson:"_id"`
	Caption          string             `json:"Caption" bson:"Caption"`
	Image_URL        string             `json:"Image_URL" bson:"Image_URL"`
	Posted_Timestamp string             `json:"Posted_Timestamp" bson:"Posted_Timestamp"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Instagram Backend API | Appointy Task")
}

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
