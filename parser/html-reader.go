package parser

import (
	"bufio"
	"cms/tag"
	"fmt"
	"log"
	"os"
)

func readLine(line string) {
	var tag tag.Tag
	var curToken, prevToken string
	for _, c := range line {
		if c == '<' {

			tag.Init()
			curToken = ""
		} else if c != '>' && c != ' ' && c != '=' && c != '"' && c != '/' {
			curToken += string(c)
		} else {
			if curToken == "!DOCTYPE" || curToken == "html" || curToken == "head" || curToken == "link" || curToken == "script" || curToken == "body" {
				return
			}
			if c == '=' {
				prevToken = curToken
				curToken = ""
			}
			if c == ' ' || c == '>' {

				if tag.Name == "" {
					tag.Name = curToken
					curToken = ""

				} else if prevToken != "" {
					if prevToken == "id" {
						tag.UID = curToken
					} else {
						tag.AddAttribute(prevToken, curToken)
					}
					prevToken, curToken = "", ""
				}
			}

		}

	}
	if tag.UID == "" {
		return
	}
	fmt.Println(tag)

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
	for scanner.Scan() {
		line := scanner.Text()
		readLine(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
