package main

/*
import (
	"cms/builder"
	"cms/parser"
)
*/

import "cms/api"

func main() {
	/*title, children := parser.ReadHTML()
	css := []string{"style", "default-style"}
	script := []string{"actions"}
	site := builder.Builder{Title: title, CSS: css, Script: script, Children: children}
	site.Build()*/
	api.CreateAPI()
}
