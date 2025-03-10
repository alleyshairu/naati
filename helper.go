package naati

import (
	"strings"
	"text/template"

	"github.com/abadojack/whatlanggo"
)

// Function to convert role enum to string
func RoleToString(language whatlanggo.Lang) string {
	switch language {
	case whatlanggo.Urd:
		return "Urdu"
	case whatlanggo.Eng:
		return "English"
	default:
		return "Unknown"
	}
}

// Function to escape latex special characters
// Like $ to \$
func LatexSpecialCharaceters(s string) string {
	return strings.ReplaceAll(s, "$", "\\$")
}

func GetFuncsMap() template.FuncMap {
	return template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"role":  RoleToString,
		"latex": LatexSpecialCharaceters,
	}
}
