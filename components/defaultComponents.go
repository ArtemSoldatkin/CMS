package components

import (
	"cms/tag"
)

// CreateLink - create lin <a></a>
func CreateLink(text, value string) *tag.Tag {
	link := tag.Tag{Name: "a", Value: text}
	link.Init()
	link.AddAttribute("href", value)
	return &link
}

// CreatePicture - create img
func CreatePicture(imgLink string) *tag.Tag {
	img := tag.Tag{Name: "img"}
	img.Init()
	img.AddAttribute("src", imgLink)
	return &img
}

// CreateText - create text
func CreateText(text string) *tag.Tag {
	tx := tag.Tag{Name: "p", Value: text}
	tx.Init()
	return &tx
}
