package main

type Project struct {
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Stylesheet  string   `json:"stylesheet"`
	URL         string   `json:"url"`
	Description string   `json:"description"`
	Writeup     []string `json:"writeup"`
	ImageHref   string   `json:"image_href"`
}
