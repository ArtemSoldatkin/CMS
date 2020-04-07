package main

import (
	"cms/tag"
	"fmt"
)

func main() {
	/*
		form := form.CreateForm("Send", "http://localhost:5000", "POST", []form.InputText{form.InputText{label: "Name", placeholder: "Enter name"}, form.InputText{label: "Description", placeholder: "Enter description"}})
		site := Builder{[]tag.Tag{*form}}
		site.build("test")*/
	span := tag.Tag{Name: "span"}
	span.Init()
	span.AddAttribute("class", "span-class")
	div := tag.Tag{Name: "div", Children: []tag.Tag{span}}
	div.Init()
	fmt.Println(div)
}
