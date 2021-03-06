package builder

import (
	"cms/parser"
	"cms/tag"
	"encoding/json"
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
	var result string
	for _, c := range b.Children {
		result += getStyles(&c)
	}
	return result
}

func (b Builder) createAction() string {
	var variables, validation, actions string
	for _, c := range b.Children {
		variables += createVariables(&c)
	}
	if variables != "" {
		variables = strings.Trim(variables, ",")
		variables = fmt.Sprintf("var %s;", variables)
	}
	for _, c := range b.Children {
		validation += createValidationFunc(&c)
	}
	for _, c := range b.Children {
		actions += createEventListeners(&c)
	}
	return fmt.Sprintf("%s%s\nwindow.onload = function(){%s\n}", variables, validation, actions)
}

// Build - make a site
func (b *Builder) Build() {
	writeToFile("index", "html", b.createDOM())
	writeToFile("style", "css", b.createStyleSheet())
	writeToFile("default-style", "css", parser.DeafultStyleToCSS())
	writeToFile("actions", "js", b.createAction())
}

// BuildToJSON - make JSON file with css / html / js
func (b *Builder) BuildToJSON() ([]byte, error) {
	result := make(map[string]string)
	result["html"] = b.createDOM()
	result["style"] = b.createStyleSheet()
	result["default-style"] = parser.DeafultStyleToCSS()
	result["actions"] = b.createAction()
	return json.Marshal(result)
}
