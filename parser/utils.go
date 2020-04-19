package parser

import (
	"cms/components"
	"cms/components/form"
	"cms/tag"
	"fmt"
)

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

func createTag(t *body) *tag.Tag {
	if t.Type == "picture" {
		return components.CreatePicture(t.Value)
	} else if t.Type == "text" {
		return components.CreateText(t.Text)
	} else if t.Type == "link" {
		return components.CreateLink(t.Text, t.Value)
	} else if t.Type == "form" {
		return form.CreateForm(t.Text, t.Value, t.Inputs)
	}
	return nil
}
