package generator

import (
	"fmt"
	"go-scaffold/internal/generator/util"
	"os"
	"path/filepath"
	"text/template"
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