package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ToPascalCase(s string) string {
	words := strings.Split(s, "_")
	c := cases.Title(language.English)
	for i, w := range words {
		words[i] = c.String(w)
	}
	return strings.Join(words, "")
}

func ToSnakeCase(s string) string {
	var result []rune
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			result = append(result, '_')
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}

func getUniqueFilePath (dir, baseName string) (string, error) {
	filePath := filepath.Join(dir, baseName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return filePath, nil
	}

	ext := filepath.Ext(baseName) // ".go"
	name := strings.TrimSuffix(baseName, ext) // "user"
	i := 1

	for {
		newName := fmt.Sprintf("%s_copy%d%s", name, i, ext)
		newPath := filepath.Join(dir, newName)
		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			return newPath, nil
		}
		i ++
		if i > 100 {
			return "", fmt.Errorf("too many duplicate files for %s", baseName)
		}
	}
}