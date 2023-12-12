package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"path/filepath"
	"strings"
)

type Template struct {
	Templ *template.Template
}

var Tmpl Template

func InitializeTemplate() {
	Tmpl.Templ = template.New("base")
	err := filepath.Walk("views", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("Error dentro de func anonima %s", err)
		}

		if path == "views" {
			return nil
		}

		if info.IsDir() {
			// Load views html
			pathStruct := strings.Split(path, "/")
			err = Tmpl.parseTemplates(pathStruct[0], pathStruct[1], "*.html")
			if err != nil {
				return err
			}

			// Load components html
			err = Tmpl.parseTemplates(pathStruct[0], pathStruct[1], "components", "*.html")
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error iterando carpetas: ", err)
	}
}

func (t *Template) parseTemplates(path ...string) error {
	htmlFiles := filepath.Join(path...)
	templ, err := t.Templ.ParseGlob(htmlFiles)
	if err != nil {
		return fmt.Errorf("Error parseando fichero: %s, %s", path, err)
	}

	t.Templ = templ // Solo actualiza t.Templ si ParseGlob fue exitoso
	return nil
}

func GetTemplates() *template.Template {
	return Tmpl.Templ
}
