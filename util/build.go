package util

import (
	"io"
	"os"
	"path/filepath"
)

type CSSData struct {
	Href   string   `json:"href"`
	Styles []string `json:"styles"`
}

func CompileCSS(pages []CSSData) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	for _, page := range pages {
		path := filepath.Join(cwd, page.Href)
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		for _, style := range page.Styles {
			path = filepath.Join(cwd, style)
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			_, err = io.Copy(file, f)
			if err != nil {
				return err
			}
			err = f.Close()
			if err != nil {
				return err
			}
		}
		err = file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
