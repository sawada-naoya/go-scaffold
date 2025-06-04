package generator

import (
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
