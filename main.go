package main

func main() {

	form := createForm("Send", []inputText{inputText{"Name", "Enter name"}, inputText{"Description", "Enter description"}, inputText{"test", "Enter test"}, inputText{}})

	html := html{[]tag{*form}}
	html.build("test")

}
