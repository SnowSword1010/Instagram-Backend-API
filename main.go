package main

import (
	"fmt"
)

type User struct {
	id       uint64 `json:"id"`
	Name     string `json:"Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type Posts struct {
	id               uint64 `json:"id"`
	Caption          string `json:"Caption"`
	Image_URL        string `json:"Image_URL"`
	Posted_Timestamp string `json:"Posted_Timestamp"`
}

func main() {
	fmt.Println("Hello")
}
