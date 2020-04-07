package tmp

/*
import (
	"cms/tag"

	"fmt"
	"log"
	"strings"
)

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

func checkTextInput(t *tag.Tag) bool {
	return t.Name == "input" && checkAttribute(t, "text")
}
*/
