package main

import "html/template"

type Project struct {
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Stylesheet  string   `json:"stylesheet"`
	URL         string   `json:"url"`
	Description string   `json:"description"`
	BodyList    []string `json:"body"`
	ImageHref   string   `json:"image_href"`
	ImageAlt    string   `json:"image_alt"`
	Body        template.HTML
}
