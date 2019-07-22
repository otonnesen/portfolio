package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/otonnesen/portfolio/util"
)

var ContentDatabase map[string]interface{}
var PageDatabase map[string]Page
var NotFound http.HandlerFunc

func main() {
	pages, err := initPageDatabase()
	if err != nil {
		log.Panicf("Error initializing page database: %v\n", err)
	}

	// Create array of CSSData structs to pass to CompileCSS
	styles := make([]util.CSSData, len(pages))
	for i, page := range pages {
		styles[i] = page.Stylesheet
	}
	err = util.CompileCSS(styles)
	if err != nil {
		log.Panicf("Error compiling CSS: %v\n", err)
	}

	ContentDatabase = make(map[string]interface{})

	err = initProjects()
	if err != nil {
		log.Panicf("Error initializing projects: %v\n", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Panicf("Error getting cwd: %v\n", err)
	}
	NotFound = util.ParseTemplate(NotFoundTemplate,
		filepath.Join(cwd, "tmpl/notfound.tmpl"))
	Root := util.ParseTemplate(RootTemplate,
		filepath.Join(cwd, "tmpl/index.tmpl"))
	Projects := util.ParseTemplate(ProjectsTemplate,
		filepath.Join(cwd, "tmpl/projects.tmpl"))
	ProjectPage := util.ParseTemplate(ProjectPageTemplate,
		filepath.Join(cwd, "tmpl/projects.tmpl"))

	http.Handle("/project/", util.LogRequest(ProjectPage))
	http.HandleFunc("/", util.LogRequest(Root))
	http.HandleFunc("/projects", util.LogRequest(Projects))
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
}

func initPageDatabase() ([]Page, error) {
	PageDatabase = make(map[string]Page)

	cwd, err := os.Getwd()
	if err != nil {
		return []Page{}, err
	}
	// Load JSON data
	f, err := os.Open(filepath.Join(cwd, "data/data.json"))
	if err != nil {
		return []Page{}, err
	}
	pages := []Page{}
	err = json.NewDecoder(f).Decode(&pages)
	if err != nil {
		return []Page{}, err
	}
	f.Close()
	for _, page := range pages {
		PageDatabase[page.Name] = page
	}

	return pages, nil
}

func initProjects() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	// Load JSON data
	f, err := os.Open(filepath.Join(cwd, "data/projects.json"))
	if err != nil {
		return err
	}

	projects := ProjectsContent{}

	err = json.NewDecoder(f).Decode(&projects)
	if err != nil {
		return err
	}
	f.Close()
	fmt.Printf("%+v\n", projects)
	ContentDatabase["projects"] = projects
	return nil
}
