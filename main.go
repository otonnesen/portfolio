package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Printf("$PORT not set, defaulting to 8080")
		port = "8080"
	}

	log.Fatal(http.ListenAndServe(":"+port, http.FileServer(http.Dir("./static/"))))
}
