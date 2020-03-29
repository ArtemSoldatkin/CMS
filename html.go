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
	return fmt.Sprintf("<!DOCTYPE html><html><head><link rel='stylesheet' href='%s.css'></head><body>%s</body></html>", fileName, result)
}

func (h html) getStyle() string {
	result := ""
	for _, c := range h.children {
		result += childrenStyleToString(c)
	}
	return result
}

func (h html) build(fileName string) {
	writeToFile(fileName, "html", h.getDOM(fileName))
	writeToFile(fileName, "css", h.getStyle())
}
