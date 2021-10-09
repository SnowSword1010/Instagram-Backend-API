package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// USER SCHEMA
type User struct {
	_id       primitive.ObjectID `json:"_id" bson:"_id"`
	UserId    uint64             `json:"User_id" bson:"User_id"`
	Name      string             `json:"Name" bson:"Name"`
	Email     string             `json:"Email" bson:"Email"`
	Password  string             `json:"Password" bson:"Password"`
	PostSlice []uint64           `json:"PostSlice" bson:"PostSlice"`
}

// POST SCHEMA
type Posts struct {
	_id              primitive.ObjectID `json:"_id" bson:"_id"`
	Email            string             `json:"Email" bson:"Email"`
	PostId           uint64             `json:"Post_id" bson:"Post_id"`
	Caption          string             `json:"Caption" bson:"Caption"`
	Image_URL        string             `json:"Image_URL" bson:"Image_URL"`
	Posted_Timestamp string             `json:"Posted_Timestamp" bson:"Posted_Timestamp"`
}

// HANDLES REQUESTS MADE TO HOME ROUTE '/'
func homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	json.NewEncoder(w).Encode("Welcome to Instagram Backend API | Appointy Task")
}

func main() {
	fmt.Println("Starting the server at port 8080")
	http.HandleFunc("/users", usersEndpoint)
	http.HandleFunc("/posts", postsEndPoint)
	http.HandleFunc("/posts/users", userPostEndPoint)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
