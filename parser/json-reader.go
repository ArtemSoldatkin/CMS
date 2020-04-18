package parser

import (
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
		//fmt.Printf(style)
	}
	return
}
