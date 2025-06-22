package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/farahsrw/manageProject/config"
	"github.com/farahsrw/manageProject/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Err loading .env file")
	}

	userSchema, err := config.LoadSchema("schemas/user.json")
	if err != nil {
		log.Fatalf("Error loading user schema: %v", err)
	}

	projectSchema, err := config.LoadSchema("schemas/project.json")
	if err != nil {
		log.Fatalf("Error loading project schema: %v", err)
	}

	db := config.InitDB()
	defer db.Close()

	jwtKey := config.GetJWTKey()

	app := &config.App{DB: db, JWTKey: jwtKey}
	router := mux.NewRouter()
	routes.SetupRoutes(router, app, userSchema, projectSchema)

	log.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
