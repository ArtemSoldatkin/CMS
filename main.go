package main

/*
import "cms/api"

func main() {
	api.CreateAPI()
}
*/

import "cms/parser"

func main() {
	parser.ReadHTMLFromFile("index")
}
