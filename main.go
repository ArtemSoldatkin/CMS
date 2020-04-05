package main

import "cms/tag"

func main() {

	form := createForm("Send", "http://localhost:5000", "POST", []inputText{inputText{label: "Name", placeholder: "Enter name"}, inputText{label: "Description", placeholder: "Enter description"}, inputText{label: "test", placeholder: "Enter test"}, inputText{}})

	html := html{[]tag.Tag{*form}}
	html.build("test")

}
