package tmp

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
