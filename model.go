package main

import "github.com/otonnesen/portfolio/util"

type Page struct {
	Name       string       `json:"name"`
	Title      string       `json:"title"`
	Stylesheet util.CSSData `json:"stylesheet"`
	Content    interface{}
}

type Project struct {
	Name        string       `json:"name"`
	Title       string       `json:"title"`
	URL         string       `json:"url"`
	Stylesheet  util.CSSData `json:"stylesheet"`
	Description string       `json:"description"`
	Writeup     []string     `json:"writeup"`
	ImageHref   string       `json:"image_href"`
}
