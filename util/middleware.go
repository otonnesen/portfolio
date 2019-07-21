package util

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// Signature for func to pass to ParseTemplate to get a HandlerFunc
type TemplateHandlerFunc func(http.ResponseWriter, *http.Request, *template.Template)

// Path to base template
var layout string = "tmpl/layout/layout.tmpl"

// Logs method, route, and processing time of a request.
func LogRequest(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h(w, r)
		end := time.Now()
		log.Printf("[%s] %q %v", r.Method, r.URL.String(), end.Sub(start))
	}
}

// Parses a TemplateHandlerFunc, returning a HandlerFunc.
func ParseTemplate(h TemplateHandlerFunc, tmpl string) http.HandlerFunc {
	t, err := template.ParseFiles(tmpl, layout)
	if err != nil {
		panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		h(w, r, t)
	}
}
