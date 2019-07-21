package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/otonnesen/portfolio/util"
)

var FakeDatabase map[string]Page
var NotFound http.HandlerFunc

func main() {
	FakeDatabase = make(map[string]Page)

	cwd, err := os.Getwd()
	if err != nil {
		log.Panicf("Error getting cwd: %v\n", err)
	}

	// Load JSON data
	f, err := os.Open(filepath.Join(cwd, "data.json"))
	if err != nil {
		log.Panicf("Error opening file: %v\n", err)
	}
	pages := []Page{}
	json.NewDecoder(f).Decode(&pages)
	f.Close()
	// Create array of CSSData structs to pass to CompileCSS
	styles := make([]util.CSSData, len(pages))
	for i, page := range pages {
		styles[i] = page.Stylesheet
	}
	err = util.CompileCSS(styles)
	if err != nil {
		log.Panicf("Error compiling CSS: %v\n", err)
	}

	for _, page := range pages {
		FakeDatabase[page.Name] = page
	}
	NotFound = util.ParseTemplate(NotFoundTemplate,
		filepath.Join(cwd, "tmpl/notfound.tmpl"))
	Root := util.ParseTemplate(RootTemplate,
		filepath.Join(cwd, "tmpl/index.tmpl"))
	Projects := util.ParseTemplate(ProjectsTemplate,
		filepath.Join(cwd, "tmpl/projects.tmpl"))
	Blog := util.ParseTemplate(BlogTemplate,
		filepath.Join(cwd, "tmpl/blog.tmpl"))
	About := util.ParseTemplate(AboutTemplate,
		filepath.Join(cwd, "tmpl/about.tmpl"))
	Contact := util.ParseTemplate(ContactTemplate,
		filepath.Join(cwd, "tmpl/contact.tmpl"))

	http.HandleFunc("/", util.LogRequest(Root))
	http.HandleFunc("/projects", util.LogRequest(Projects))
	http.HandleFunc("/blog", util.LogRequest(Blog))
	http.HandleFunc("/about", util.LogRequest(About))
	http.HandleFunc("/contact", util.LogRequest(Contact))
	static := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", static)
	http.ListenAndServe(":8080", nil)
}
