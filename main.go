package main

func main() {

	span := tag{name: "span", attributes: []string{"class='span'"}}
	span.Init()
	button := tag{name: "button", attributes: []string{"class='button'"}, value: "ok"}
	button.Init()
	img := tag{name: "img", attributes: []string{"class='img'"}}
	img.Init()

	div := tag{name: "div", attributes: []string{"id='id'", "class='class'"}, children: []tag{span, button, img}}
	div.Init()

	div.appendChildren([]tag{span, span})
	div.removeChildren([]tag{img})

	div.switchChildren(button, 0)

	html := html{[]tag{div}}
	writeToFile("test", html.String())

}
