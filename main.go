package main

import (
	"cms/builder"
	"cms/form"
	"cms/tag"
)

func main() {
	formInputs := []form.InputText{form.InputText{Label: "Name", Placeholder: "Enter name", ValueName: "name", NotValidMsg: "Name is empty"}, form.InputText{Label: "Description", Placeholder: "Enter description", ValueName: "description", NotValidMsg: "Description is empty"}}
	form := form.CreateForm("Send", "http://localhost:5000", "POST", formInputs)
	//form.AddStyle("color", "red")
	//form.AddStyle("background", "green")
	form.AddStyle("font-size", "24px")

	title := "Test"
	css := []string{"test"}
	script := []string{"test"}
	site := builder.Builder{Title: title, CSS: css, Script: script, Children: []tag.Tag{*form}}
	site.Build()

}
