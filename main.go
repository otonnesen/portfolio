package main

import (
	"log"
	"net/http"
	"os"
)

var logger *log.Logger

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Printf("$PORT not set, defaulting to 8080")
		port = "8080"
	}

	logFile, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer logFile.Close()

	logger = log.New(logFile, "", log.LstdFlags)

	m := &middleware{http.FileServer(http.Dir("./static/"))}
	log.Fatal(http.ListenAndServe(":"+port, m))
}

type middleware struct {
	h http.Handler
}

func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger.Printf("[%s]: %s %q", r.RemoteAddr, r.Method, r.URL.String())
	m.h.ServeHTTP(w, r)
}
