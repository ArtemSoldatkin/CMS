package parser

import (
	"bytes"
	HTMLTag "cms/tag"
	"fmt"
	"os"
	"strings"
)

// Tree node
type node struct {
	tag               *HTMLTag.Tag
	isClosed, isChild bool
}

func (n *node) close() {
	n.isClosed = true
}

// Html tree
type tree struct {
	tree []*node
}

func (t *tree) String() string {
	var parent []string
	for _, n := range t.tree {
		if !n.isChild {
			parent = append(parent, n.tag.String())
		}
	}
	return strings.Join(parent, "")
}

func (t *tree) addAttribute(key, value string) {
	if len(t.tree) > 0 {
		t.tree[len(t.tree)-1].tag.AddAttribute(key, value)
	}
}

func (t *tree) setValue(value string) {
	if len(t.tree) > 0 {
		t.tree[len(t.tree)-1].tag.Value += value
	}
}

func (t *tree) backward() *node {
	for i := len(t.tree) - 1; i >= 0; i-- {
		if !t.tree[i].isClosed {
			return t.tree[i]
		}
	}
	return nil
}

func (t *tree) addNode(n *node) {
	parentNode := t.backward()
	if parentNode != nil {
		n.isChild = true
		parentNode.tag.AppendChild(n.tag)
	}
	t.tree = append(t.tree, n)
}

func (t *tree) closeNode() {
	node := t.backward()
	if node != nil {
		node.close()
	}
}

// Html reader

type reader struct {
	html, tag       string
	pos             int
	isQuot, isValue bool
	nodeTree        tree
}

func (r *reader) init() {
	r.tag = r.getTag()
}

func (r *reader) String() string {
	return r.nodeTree.String()
}

func (r *reader) getTag() string {
	tag := ""
	for r.pos < len(r.html) {
		char := string(r.html[r.pos])
		r.pos++
		if r.isValue {
			if char != "<" {
				tag += char
				continue
			} else {
				r.pos--
				return tag
			}
		}

		if char == ">" {
			return char
		}

		if char == "<" {
			if tag != "" {
				r.pos--
				return tag
			}
			return char
		}

		if char == "\"" || char == "'" {
			r.isQuot = !r.isQuot
			if !r.isQuot {
				return tag
			}
			continue
		}

		if !r.isQuot {
			if char == "/" {
				return char
			}
			if char == " " || char == "\t" || char == "=" || char == "\n" {
				if tag != "" {
					return tag
				}
				continue
			}
		}

		tag += char
	}
	return tag
}

func (r *reader) read() {
	for r.pos < len(r.html) {
		if r.tag == "<" {
			tag := r.getTag()
			r.isValue = false
			if tag != "/" {
				htmlTag := HTMLTag.Tag{Name: tag}
				htmlTag.Init()
				r.nodeTree.addNode(&node{tag: &htmlTag, isClosed: false})
			} else {
				r.nodeTree.closeNode()
			}
		} else if r.tag == "/" {
			tag := r.getTag()
			if tag == ">" {
				r.nodeTree.closeNode()
			}
		} else if r.tag == ">" {
			r.isValue = true
			tag := r.getTag()
			if tag != "<" {
				r.nodeTree.setValue(tag)
				r.isValue = false
			} else {
				r.isValue = false
				continue
			}
		} else if !r.isValue {
			r.nodeTree.addAttribute(r.tag, r.getTag())
		}
		r.tag = r.getTag()
	}
}

// ReadHTMLFromFile - read and create html from html file
func ReadHTMLFromFile(fileName string) {
	rootDir, _ := os.Getwd()
	file, err := os.Open(fmt.Sprintf("%s/parser/test-data/%s.html", rootDir, fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var buf bytes.Buffer
	_, err = buf.ReadFrom(file)
	if err != nil {
		panic(err)
	}
	html := buf.String()
	r := reader{html: html}
	r.init()
	r.read()
	fmt.Println(r.String())
}
