package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func computeHash() {

}

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
				user.Name = r.Form["Name"][0]
				user.Email = r.Form["Email"][0]
				// ENCRYPTING AND STORING PASSWORD
				user.Password = givePwdHash(r.Form["Password"][0])
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
	}
}

func main() {
	fmt.Println("Starting the server at port 8080")

	// client, _ = mongo.Connect(ctx, "mongodb://localhost:27017")
	http.HandleFunc("/users", usersEndpoint)
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
