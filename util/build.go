package util

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

func CompileCSS(pages []string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	for _, page := range pages {
		path := filepath.Join(cwd, "static", "build", page+".css")
		log.Printf("Compiling style %s...\n", path)
		file, err := os.Create(path)
		if err != nil {
			return err
		}
		styles := []string{
			filepath.Join(cwd, "static", "layout.css"),
			filepath.Join(cwd, "static", page+".css"),
		}
		for _, style := range styles {
			f, err := os.Open(style)
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
