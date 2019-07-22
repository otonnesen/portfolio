package main

import "github.com/otonnesen/portfolio/util"

type Page struct {
	Name       string       `json:"name"`
	Title      string       `json:"title"`
	URL        string       `json:"url,omitempty"`
	Stylesheet util.CSSData `json:"stylesheet"`
	Content    interface{}
}

type ProjectListEntry struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type ProjectsContent struct {
	Projects []ProjectListEntry `json:"projects"`
}
