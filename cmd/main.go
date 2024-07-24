package main

import (
	"log"

	"github.com/Azpect3120/CloudDatabaseSpawner/internal/scripts/users"
)

const POSTGRES_PASSWORD = "Panther4487!!!!"

func main() {
	// if err := scripts.CreateUser("Hayden", "root", POSTGRES_PASSWORD); err != nil {
	// 	log.Fatalf("Failed to create user: %v", err)
	// }
	if err := scripts.CreateUserNoScript("Hayden", "root", POSTGRES_PASSWORD); err != nil {
		log.Fatalf("Failed to create user: %v", err)
	}
}
