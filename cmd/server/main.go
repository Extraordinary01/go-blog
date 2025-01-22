package main

import (
	"fmt"
	"go-blog/api"
	"go-blog/database"
	"log"
	"os"
)

func main () {
	log.Printf("Starting server at port 8080")

	db, err := database.StartDB()
	if err != nil {
		log.Fatalf("could not start database: %v", err)
	}
	defer db.Close()

	router := api.InitAPI(db)

	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
	// Not tested, should test tomorrow, after succesfull testing, need to write controller for blogs.
}