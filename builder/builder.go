package builder

import (
	"cms/tag"
	"fmt"
	"strings"
)

// Builder - struct to create site.
type Builder struct {
	Title    string
	Script   []string
	CSS      []string
	Children []tag.Tag
}

func (b Builder) createDOM() string {
	var params, children []string
	for _, p := range b.CSS {
		params = append(params, createCSSLink(p))
	}
	for _, p := range b.Script {
		params = append(params, createScriptLink(p))
	}
	head := createHead(b.Title, params...)
	for _, c := range b.Children {
		children = append(children, c.String())
	}
	body := createBody(&children)
	return fmt.Sprintf("<!DOCTYPE html>\n<html>\n%s\n%s\n</html>", head, body)
}

func (b Builder) createStyleSheet() string {
	var style []string
	for _, c := range b.Children {
		style = append(style, c.StyleToString())
	}
	return strings.Join(style, "\n")
}

func (b Builder) createAction() string {
	var variables, actions string
	for _, c := range b.Children {
		variables += createVariables(&c)
	}
	if variables != "" {
		variables = strings.Trim(variables, ",")
		variables = fmt.Sprintf("var %s;", variables)
	}
	for _, c := range b.Children {
		actions += createEventListeners(&c)
	}
	return fmt.Sprintf("%s\nwindow.onload = function(){%s\n}", variables, actions)
}

// Build - make a site
func (b *Builder) Build() {
	writeToFile("test", "html", b.createDOM())
	writeToFile("test", "css", b.createStyleSheet())
	writeToFile("test", "js", b.createAction())
}
