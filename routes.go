package main

import (
	"html/template"
	"net/http"
	"strings"
)

func RootTemplate(w http.ResponseWriter, r *http.Request,
	t *template.Template) {
	url := r.URL.String()
	data := PageDatabase["home"]
	// Returns 404 if URL contains anything other than "/"
	if url != "/" {
		NotFound(w, r)
		return
	}
	t.Execute(w, data)
}

func ProjectsTemplate(w http.ResponseWriter, r *http.Request,
	t *template.Template) {
	url := r.URL.String()
	data := PageDatabase["projects"]
	if url != "/projects" {
		NotFound(w, r)
		return
	}
	p := struct {
		Projects []Project
	}{Projects}
	data.Content = p
	t.Execute(w, data)
}

func ProjectPageTemplate(w http.ResponseWriter, r *http.Request,
	t *template.Template) {
	url := r.URL.String()
	name := strings.Split(url, "/")[2]
	if url != "/project/"+name /* TODO: or if project doesn't exist */ {
		NotFound(w, r)
		return
	}
	data := ProjectDatabase[name]
	t.Execute(w, data)
}

func NotFoundTemplate(w http.ResponseWriter, r *http.Request,
	t *template.Template) {
	data := PageDatabase["notfound"]

	t.Execute(w, data)
}
