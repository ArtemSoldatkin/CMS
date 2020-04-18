package parser

import "fmt"

func getStyle(t *tagStyle, name string) string {
	var style, result string
	for k, v := range t.Style {
		style += fmt.Sprintf("\t%s: %s;\n", k, v)
	}
	if name == "" {
		name = t.Name
	} else {
		name = fmt.Sprintf("%s %s", name, t.Name)
	}
	for _, c := range t.Children {
		result += getStyle(&c, t.Name)
	}
	return fmt.Sprintf("%s\n %s {\n%s}\n", result, name, style)
}
