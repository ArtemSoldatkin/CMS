package main

func main() {

	form := createForm("Send", []inputText{inputText{label: "Name", placeholder: "Enter name"}, inputText{label: "Description", placeholder: "Enter description"}, inputText{label: "test", placeholder: "Enter test"}, inputText{}})

	html := html{[]tag{*form}}
	html.build("test")

}
