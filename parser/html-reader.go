package parser

import (
	"bufio"
	"cms/tag"
	"fmt"
	"log"
	"os"
)

func getToken(line string, position *int, quot *bool) (token string) {
	for _, c := range line[*position:] {
		*position++
		if c == '\'' || c == '"' {
			*quot = !*quot
			continue
		}
		if !*quot && (c == ' ' || c == '=') {
			return
		}
		token += string(c)
		if c == '>' || c == '<' || (c == '/' && !*quot) {
			return
		}
	}
	return
}

func readLine(line string, t *[]tag.Tag) {
	var currentTag tag.Tag
	var position int
	var quot, isValue bool
	token := getToken(line, &position, &quot)
	for position < len(line) {
		if token == "<" {
			token = getToken(line, &position, &quot)
			if token == "/" {
				isValue, quot = false, false
			} else {
				currentTag = tag.Tag{Name: token}
				currentTag.Init()
				*t = append(*t, currentTag)
			}
		} else if token == "/" && getToken(line, &position, &quot) == ">" {
			// pass
		} else {
			if isValue {
				currentTag.Value = getToken(line, &position, &quot)
			} else {
				if token == "id" {
					currentTag.UID = getToken(line, &position, &quot)
				} else if token == "value" {
					currentTag.Value = getToken(line, &position, &quot)
				} else {
					currentTag.AddAttribute(token, getToken(line, &position, &quot))
				}
			}
		}
		token = getToken(line, &position, &quot)
		if token == ">" {
			isValue, quot = true, true
		}
	}
}

// ReadHTMLFromFile - read and create html from html file
func ReadHTMLFromFile(fileName string) {
	rootDir, _ := os.Getwd()
	html, err := os.Open(fmt.Sprintf("%s/parser/test-data/%s.html", rootDir, fileName))
	if err != nil {
		panic(err)
	}
	defer html.Close()

	scanner := bufio.NewScanner(html)
	var t []tag.Tag
	for scanner.Scan() {
		line := scanner.Text()
		readLine(line, &t)
	}
	fmt.Println(t)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
