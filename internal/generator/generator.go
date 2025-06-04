package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/sawada-naoya/go-scaffold/internal/generator/util"
)

func Generator(name string, layers []string) error {
	structName := util.ToPascalCase(name)
	fileName := util.ToSnakeCase(name)

	for _, layer := range layers {
		dir := filepath.Join("internal", layer)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}

		templPath := filepath.Join("internal", "template", layer+".tmpl")
		tmpl, err := template.ParseFiles(templPath)
		if err != nil {
			return err
		}

		filePath := filepath.Join(dir, fileName)
		f, err := os.Create(filePath)
		if err != nil  {
			return err
		}

		defer f.Close()

		data := map[string]string{"StructName": structName}
		if err := tmpl.Execute(f, data); err != nil {
			return err
		}
		fmt.Printf("Created: %s\n", filePath)
	}
	return nil
}