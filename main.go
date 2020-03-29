package main

func main() {

	span := tag{name: "span", style: []string{"width:100px", "height:100px", "background:red"}, value: "some text"}
	span.Init()

	buttonAction := action{Actions.Click, "console.log('Hello World')"}
	button := tag{name: "button", style: []string{"width:100px", "height:100px"}, value: "ok", action: []action{buttonAction}}
	button.Init()

	img := tag{name: "img", style: []string{"width:100px", "height:100px"}}
	img.Init()

	inputName := tag{name: "input", attributes: []string{"type='text'"}}
	inputName.Init()

	inputDescription := tag{name: "input", attributes: []string{"type='text'"}}
	inputDescription.Init()

	inputSubmit := tag{name: "input", value: "send", attributes: []string{"type='submit'"}}
	inputSubmit.Init()

	formAction := action{Actions.Submit, "alert('ok')"}
	form := tag{name: "form", action: []action{formAction}, children: []tag{inputName, inputDescription, inputSubmit}}
	form.Init()

	div := tag{name: "div", style: []string{"width:1000px", "height:1000px", "background:green"}, children: []tag{span, button, img, form}}
	div.Init()

	/*
		div.appendChildren([]tag{span, span})
		div.removeChildren([]tag{img})

		div.switchChildren(button, 0)
	*/
	html := html{[]tag{div}}
	html.build("test")

}
