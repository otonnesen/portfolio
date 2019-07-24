package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/otonnesen/portfolio/util"
)

var Projects []Project
var ProjectDatabase map[string]Project
var NotFound http.HandlerFunc

func main() {
	// Create array of CSSData structs to pass to CompileCSS
	styles := []string{
		"home",
		"projects",
		"projectpage",
		"notfound",
	}
	err := util.CompileCSS(styles)
	if err != nil {
		log.Panicf("Error compiling CSS: %v\n", err)
	}

	ProjectDatabase = make(map[string]Project)

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
		filepath.Join(cwd, "tmpl/projectpage.tmpl"))

	http.Handle("/project/", util.LogRequest(ProjectPage))
	http.HandleFunc("/", util.LogRequest(Root))
	http.HandleFunc("/projects", util.LogRequest(Projects))
	http.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":8080", nil)
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

	// projects := []Project{}

	err = json.NewDecoder(f).Decode(&Projects)
	if err != nil {
		return err
	}
	f.Close()
	for _, project := range Projects {
		ProjectDatabase[project.Name] = project
	}
	return nil
}
