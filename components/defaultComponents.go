package components

import (
	"cms/tag"
	"fmt"
)

// LinkAttributes - link parameters
type LinkAttributes struct {
	Text, LinkValue, LinkType string
}

// CreateLink - create link
func CreateLink(attr LinkAttributes) *tag.Tag {
	link := tag.Tag{Name: "a", Value: attr.Text}
	link.Init()
	var href string
	if attr.LinkType == "tel" {
		href = fmt.Sprintf("tel:%s", attr.LinkValue)
	} else if attr.LinkType == "mail" {
		href = fmt.Sprintf("mailto:", attr.LinkValue)
	} else {
		href = attr.LinkValue
	}
	link.AddAttribute("href", href)
	return &link
}

// CreateImg - create img
func CreateImg(imgLink string) *tag.Tag {
	img := tag.Tag{Name: "img"}
	img.Init()
	img.AddAttribute("src", imgLink)
	return &img
}

// CreateText - create text
func CreateText(text string) *tag.Tag {
	return &tag.Tag{Name: "p", Value: text}
}
