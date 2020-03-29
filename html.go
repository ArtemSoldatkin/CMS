package main

import "fmt"

type html struct {
	children []tag
}

func (h html) getDOM(fileName string) string {
	result := ""
	for _, c := range h.children {
		result += childrenToString(c)
	}
	return fmt.Sprintf("<!DOCTYPE html><html><head><link rel='stylesheet' href='%s.css'><script type='text/javascript' src='%s.js'></script></head><body>%s</body></html>", fileName, fileName, result)
}

func (h html) getStyle() string {
	result := ""
	for _, c := range h.children {
		result += childrenStyleToString(c)
	}
	return result
}

func (h html) getAction() string {
	result := ""
	for _, c := range h.children {
		result += childrenActionToString(c)
	}
	return fmt.Sprintf("window.onload=function(){%s}", result)
}

func (h html) build(fileName string) {
	writeToFile(fileName, "html", h.getDOM(fileName))
	writeToFile(fileName, "css", h.getStyle())
	writeToFile(fileName, "js", h.getAction())
}
