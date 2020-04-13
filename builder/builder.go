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

// Build - make a site
func (b *Builder) Build() {
	writeToFile("test", "html", b.createDOM())
	writeToFile("test", "css", b.createStyleSheet())
}

/*
import (
	"cms/tag"
	"fmt"
	"strings"
)

// Builder - site builder
type Builder struct {
	Children []tag.Tag
}

func (h Builder) getDOM(fileName string) string {
	result := ""
	for _, c := range h.children {
		result += childrenToString(c)
	}
	return fmt.Sprintf("<!DOCTYPE html><html><head><link rel='stylesheet' href='%s.css'><script type='text/javascript' src='%s.js'></script></head><body>%s</body></html>", fileName, fileName, result)
}

func (h Builder) getStyle() string {
	result := ""
	for _, c := range h.children {
		result += childrenStyleToString(c)
	}
	return result
}

/*
func (h html) getAction() string {
	result := ""
	for _, c := range h.children {
		result += childrenActionToString(c)
	}
	return fmt.Sprintf("window.onload=function(){%s}", result)
}


func (h Builder) getAction() string {
	var variables, setVariables, result string
	for _, c := range h.children {
		variables += strings.Trim(createVariables(&c), ",")
		setVariables += createOnChange(&c)
		result += childrenActionToString(c)
	}
	if variables != "" {
		variables = fmt.Sprintf("var %s\n", variables)
	}
	return fmt.Sprintf("%swindow.onload=function(){\n%s%s}", variables, setVariables, result)
}

// Build - build site
func (h Builder) Build(fileName string) {
	writeToFile(fileName, "html", h.getDOM(fileName))
	writeToFile(fileName, "css", h.getStyle())
	writeToFile(fileName, "js", h.getAction())
}
*/
