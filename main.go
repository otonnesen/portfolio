package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	fmt.Printf("Running...\n")

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

	logger := log.New(logFile, "", log.LstdFlags)

	m := &middleware{logger, http.FileServer(exclDirFs{http.Dir("./static/")})}
	log.Fatal(http.ListenAndServe(":"+port, m))
}

type middleware struct {
	l *log.Logger
	h http.Handler
}

func (m *middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.l.Printf("[%s]: %s %q", r.RemoteAddr, r.Method, r.URL.String())
	m.h.ServeHTTP(w, r)
}

type exclDirFs struct {
	fs http.FileSystem
}

// Wrapper for http.Dir to exclude directories
func (fs exclDirFs) Open(path string) (http.File, error) {
	f, err := fs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := fs.fs.Open(index); err != nil {
			if e := f.Close(); e != nil {
				return nil, e
			}

			return nil, err
		}
	}

	return f, nil
}
