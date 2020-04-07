package form

/*
import (
	"cms/tag"
	"cms/utils"
	"fmt"
	"strings"
)

// InputText - text field for form
type InputText struct {
	Label, Placeholder, ValueName, Validation string
}

func createTextField(params inputText) []tag.Tag {
	valueName := params.valueName
	if valueName == "" {
		valueName = params.label
	}
	input := tag.Tag{Name: "input", Attributes: []string{"type=\"text\""}, ValueName: valueName}
	input.Init()
	if params.placeholder != "" {
		input.AddAttribute(fmt.Sprintf("placeholder=\"%s\"", params.placeholder))
	}
	if params.label != "" {
		label := tag.Tag{Name: "label", Value: params.label, Attributes: []string{fmt.Sprintf("for=%s", input.UID)}}
		label.Init()
		return []tag.Tag{label, input}
	}
	return []tag.Tag{input}
}

func onChange(input *tag.Tag) string {
	return fmt.Sprintf("document.getElementById(\"%s\").addEventListener(\"change\",function(e){%s=e.target.value})\n", input.UID, strings.ReplaceAll(input.UID, "-", ""))
}

func createValidation(t *tag.Tag) string {
	if t.Validation == "" {
		return fmt.Sprintf("!%s", utils.getVariable(t))
	}
	return t.Validation
}
*/
