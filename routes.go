package main

import (
	"html/template"
	"net/http"
	"strings"
)

func RootTemplate(w http.ResponseWriter, r *http.Request,
	t *template.Template) {
	url := r.URL.String()
	data := struct {
		Name       string
		Title      string
		Stylesheet string
	}{
		Name:       "home",
		Title:      "Home",
		Stylesheet: "static/build/home.css",
	}
	if url != "/" {
		NotFound(w, r)
		return
	}
	t.Execute(w, data)
}

func ProjectsTemplate(w http.ResponseWriter, r *http.Request,
	t *template.Template) {
	url := r.URL.String()
	data := struct {
		Name        string
		Title       string
		Stylesheet  string
		ProjectList []Project
	}{
		Name:        "projects",
		Title:       "Projects",
		Stylesheet:  "static/build/projects.css",
		ProjectList: Projects,
	}
	if url != "/projects" {
		NotFound(w, r)
		return
	}
	t.Execute(w, data)
}

func ProjectPageTemplate(w http.ResponseWriter, r *http.Request,
	t *template.Template) {
	url := r.URL.String()
	name := strings.Split(url, "/")[2]
	if _, ok := ProjectDatabase[name]; !ok || url != "/project/"+name {
		NotFound(w, r)
		return
	}
	data := ProjectDatabase[name]
	t.Execute(w, data)
}

func NotFoundTemplate(w http.ResponseWriter, r *http.Request,
	t *template.Template) {
	data := struct {
		Name       string
		Title      string
		Stylesheet string
	}{
		Name:       "notfound",
		Title:      "404: Not Found",
		Stylesheet: "static/build/notfound.css",
	}

	t.Execute(w, data)
}
