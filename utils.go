package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"strings"
)

type commonError struct {
	message string
}

func (e *commonError) Error() string {
	return e.message
}

func childrenToString(n tag) string {
	result := ""
	for _, c := range n.children {
		result += childrenToString(c)
	}
	if n.name == "input" {
		return fmt.Sprintf("<%s id=\"%s\" %s %s/>", n.name, n.uid, strings.Join(n.attributes, " "), n.value)
	}
	return fmt.Sprintf("<%s id='%s' %s>%s%s</%s>", n.name, n.uid, strings.Join(n.attributes, " "), n.value, result, n.name)
}

func childrenStyleToString(t tag) string {
	result := ""
	for _, c := range t.children {
		result += childrenStyleToString(c)
	}
	return fmt.Sprintf("%s#%s {%s;}\n", result, t.uid, strings.Join(t.style, "; "))
}

func childrenActionToString(t tag) string {
	result := ""
	if t.name == "form" {
		if len(t.action) > 0 {
			t.action[0].addEvent(getChildrenValues(&t))
		}
	}
	for _, c := range t.children {
		result += childrenActionToString(c)
	}
	for _, a := range t.action {
		result += fmt.Sprintf("%s\n", a.toString(t.uid))
	}
	return result
}

func generateUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("id%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func findChildPos(t *tag, child tag) (int, error) {
	for i, c := range t.children {
		if c.uid == child.uid {
			return i, nil
		}
	}
	return -1, &commonError{"child is not found"}
}

func checkAttribute(t *tag, text string) bool {
	for _, attr := range t.attributes {
		if strings.Contains(attr, text) {
			return true
		}
	}
	return false
}

func getChildrenValues(t *tag) string {
	result := ""
	for _, c := range t.children {
		if c.name != "input" && !checkAttribute(&c, "text") {
			continue
		}
		result += getChildrenValues(&c)
	}
	if t.name != "input" && !checkAttribute(t, "text") {
		return result
	}
	return fmt.Sprintf("%salert(document.getElementById('%s').value)\n", result, t.uid)
}
