package parser

import (
	"cms/components/form"
	"cms/tag"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readJSON(fileName string) []byte {

	rootDir, _ := os.Getwd()
	jsonFile, err := os.Open(fmt.Sprintf("%s/parser/%s.json", rootDir, fileName))
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	return byteValue

}

type style struct {
	CSS []tagStyle `json:"css"`
}

type tagStyle struct {
	Name     string            `json:"name"`
	Style    map[string]string `json:"style"`
	Children []tagStyle        `json:"children"`
}

// ReadDefaultStyle - read default style from default-style.json
func readDefaultStyle() style {
	var defaultStyle style
	byteValue := readJSON("default-style")
	err := json.Unmarshal(byteValue, &defaultStyle)
	if err != nil {
		panic(err)
	}
	return defaultStyle
}

// DeafultStyleToCSS - convert default style to css
func DeafultStyleToCSS() (style string) {
	defaultStyle := readDefaultStyle()
	for _, t := range defaultStyle.CSS {
		style += fmt.Sprintf("%s\n", getStyle(&t, ""))
	}
	return
}

type html struct {
	Title string `json:"title"`
	Body  []body `json:"body"`
}

type body struct {
	Type   string           `json:"type"`   // [text / form / link / picture]
	Text   string           `json:"text"`   // [text / form / link]
	Value  string           `json:"value"`  // [link / picture]
	Inputs []form.InputText `json:"inputs"` // [form]
}

// ReadHTML - read html from json file.
func ReadHTML(byteValue []byte) (string, []tag.Tag) {
	var dom html
	err := json.Unmarshal(byteValue, &dom)
	if err != nil {
		panic(err)
	}
	var domNodes []tag.Tag
	for _, t := range dom.Body {
		tg := createTag(&t)
		if tg != nil {
			domNodes = append(domNodes, *tg)
		}
	}
	return dom.Title, domNodes
}

// ReadHTMLFromFile - read html json from file
func ReadHTMLFromFile() (string, []tag.Tag) {
	byteValue := readJSON("test-data/html")
	return ReadHTML(byteValue)
}
