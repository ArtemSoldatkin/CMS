package main

import (
	"cms/form"
	"fmt"
)

func main() {
	/*
		form := form.CreateForm("Send", "http://localhost:5000", "POST", []form.InputText{form.InputText{label: "Name", placeholder: "Enter name"}, form.InputText{label: "Description", placeholder: "Enter description"}})
		site := Builder{[]tag.Tag{*form}}
		site.build("test")*/

	formInputs := []form.InputText{form.InputText{Label: "Name", Placeholder: "Enter name", ValueName: "name"}, form.InputText{Label: "Description", Placeholder: "Enter description", ValueName: "description"}}
	form := form.CreateForm("Send", "url", "POST", formInputs)

	fmt.Println(form)

}
