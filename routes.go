package main

import (
	"html/template"
	"net/http"

	"github.com/otonnesen/portfolio/util"
)

type Page struct {
	Name            string          `json:"name"`
	Title           string          `json:"title"`
	URL             string          `json:"url,omitempty"`
	Stylesheet      util.CSSData    `json:"stylesheet"`
	ProjectsContent ProjectsContent `json:"projects_content,omitempty"`
	IndexContent    IndexContent    `json:"index_content,omitempty"`
	AboutContent    IndexContent    `json:"about_content,omitempty"`
}

type ProjectListEntry struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type ProjectsContent struct {
	Projects []ProjectListEntry `json:"projects"`
}

type IndexContent struct {
}

type BlogContent struct {
}

type AboutContent struct {
}

type ContactContent struct {
}

func RootTemplate(w http.ResponseWriter, r *http.Request, t *template.Template) {
	url := r.URL.String()
	// Returns 404 if URL contains anything other than "/"
	data := FakeDatabase["home"]
	if url != "/" {
		NotFound(w, r)
		return
	}
	t.Execute(w, data)
}

func ProjectsTemplate(w http.ResponseWriter, r *http.Request, t *template.Template) {
	url := r.URL.String()
	data := FakeDatabase["projects"]
	if url != "/projects" {
		NotFound(w, r)
		return
	}
	t.Execute(w, data)
}

func BlogTemplate(w http.ResponseWriter, r *http.Request, t *template.Template) {
	url := r.URL.String()
	data := FakeDatabase["blog"]
	if url != "/blog" {
		NotFound(w, r)
		return
	}
	t.Execute(w, data)
}

func AboutTemplate(w http.ResponseWriter, r *http.Request, t *template.Template) {
	url := r.URL.String()
	data := FakeDatabase["about"]
	if url != "/about" {
		NotFound(w, r)
		return
	}
	t.Execute(w, data)
}

func ContactTemplate(w http.ResponseWriter, r *http.Request, t *template.Template) {
	url := r.URL.String()
	data := FakeDatabase["contact"]
	if url != "/contact" {
		NotFound(w, r)
		return
	}
	t.Execute(w, data)
}

func NotFoundTemplate(w http.ResponseWriter, r *http.Request, t *template.Template) {
	data := FakeDatabase["notfound"]
	t.Execute(w, data)
}
