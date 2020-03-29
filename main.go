package main

func main() {

	span := tag{name: "span", style: []string{"width:100px", "height:100px", "background:red"}, value: "some text"}
	span.Init()
	button := tag{name: "button", style: []string{"width:100px", "height:100px"}, value: "ok"}
	button.Init()
	img := tag{name: "img", style: []string{"width:100px", "height:100px"}}
	img.Init()

	div := tag{name: "div", style: []string{"width:1000px", "height:1000px", "background:green"}, children: []tag{span, button, img}}
	div.Init()

	/*
		div.appendChildren([]tag{span, span})
		div.removeChildren([]tag{img})

		div.switchChildren(button, 0)
	*/
	html := html{[]tag{div}}
	html.build("test")

}
