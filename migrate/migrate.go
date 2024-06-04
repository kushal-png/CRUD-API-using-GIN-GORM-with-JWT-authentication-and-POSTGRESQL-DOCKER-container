package main

import (
	"fmt"
	"jwt-golang/initializers"
	"jwt-golang/models"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalf("could not load environment variables")
	}

	initializers.ConnectDB(&config)
}

func main() {
	err := initializers.DB.AutoMigrate(&models.User{},&models.Post{})
	if err != nil {
		fmt.Println("Error in Migration")
		return
	}
	fmt.Println("Migration Completed")
}
