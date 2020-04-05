package main

import (
	"cms/tag"
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

func childrenToString(n tag.Tag) string {
	result := ""
	for _, c := range n.Children {
		result += childrenToString(c)
	}
	if n.Name == "input" {
		return fmt.Sprintf("<%s id=\"%s\" %s %s/>", n.Name, n.UID, strings.Join(n.Attributes, " "), n.Value)
	}
	return fmt.Sprintf("<%s id='%s' %s>%s%s</%s>", n.Name, n.UID, strings.Join(n.Attributes, " "), n.Value, result, n.Name)
}

func childrenStyleToString(t tag.Tag) string {
	result := ""
	for _, c := range t.Children {
		result += childrenStyleToString(c)
	}
	return fmt.Sprintf("%s#%s {%s;}\n", result, t.UID, strings.Join(t.Style, "; "))
}

func childrenActionToString(t tag.Tag) string {
	result := ""
	if t.Name == "form" {
		if len(t.Actions) > 0 {
			t.Actions[0].AddEvent(getChildrenValues(&t))
		}
	}
	for _, c := range t.Children {
		result += childrenActionToString(c)
	}
	for _, a := range t.Actions {
		result += fmt.Sprintf("%s", a.ToString(t.UID))
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

func findChildPos(t *tag.Tag, child tag.Tag) (int, error) {
	for i, c := range t.Children {
		if c.UID == child.UID {
			return i, nil
		}
	}
	return -1, &commonError{"child is not found"}
}

func checkAttribute(t *tag.Tag, text string) bool {
	for _, attr := range t.Attributes {
		if strings.Contains(attr, text) {
			return true
		}
	}
	return false
}

func getChildrenValues(t *tag.Tag) string {
	result := ""
	if t.Name != "input" && !checkAttribute(t, "text") {
		return result
	}
	for _, c := range t.Children {
		result += getChildrenValues(&c)
	}
	return fmt.Sprintf("%salert(document.getElementById('%s').value)\n", result, t.UID)
}

func createVariables(t *tag.Tag) string {
	var result string
	for _, c := range t.Children {
		if !checkTextInput(&c) {
			continue
		}
		result += createVariables(&c)
	}
	if !checkTextInput(t) {
		return result
	}
	return fmt.Sprintf("%s,%s", result, strings.ReplaceAll(t.UID, "-", ""))
}

func createOnChange(t *tag.Tag) string {
	var result string
	for _, c := range t.Children {
		if !checkTextInput(&c) {
			continue
		}
		result += onChange(&c)
	}
	if !checkTextInput(t) {
		return result
	}
	return fmt.Sprintf("%s\n%s", result, onChange(t))
}

func getVariable(t *tag.Tag) string {
	return strings.ReplaceAll(t.UID, "-", "")
}
