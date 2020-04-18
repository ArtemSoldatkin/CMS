package form

import (
	"cms/tag"
	"fmt"
)

// InputText - text field for form
type InputText struct {
	Label, Placeholder, ValueName, Validation, NotValidMsg string
}

func createTextField(params InputText) []tag.Tag {
	var result []tag.Tag
	input := tag.Tag{Name: "input", Value: "", ValueName: params.ValueName}
	input.Init()
	input.AddAttribute("type", "text")
	if params.Placeholder != "" {
		input.AddAttribute("placeholder", params.Placeholder)
	}

	result = append(result, input)

	span := tag.Tag{Name: "span", Value: params.NotValidMsg}
	span.Init()
	span.AddStyle("display", "inline")
	span.AddStyle("color", "red")

	event := fmt.Sprintf("%s=e.target.value;\n%sIsValid(\"%s\");\n", input.UID, input.UID, span.UID)
	input.AddAction("change", event)

	event = fmt.Sprintf("%sIsValid(\"%s\");\n", input.UID, span.UID)
	input.AddAction("blur", event)

	if params.Label != "" {
		label := tag.Tag{Name: "label", Value: params.Label}
		label.Init()
		label.AddAttribute("for", input.UID)
		return []tag.Tag{label, input, span}
	}

	return []tag.Tag{input, span}
}

func createInputValidation(t *tag.Tag) string {
	if t.Name != "input" {
		return ""
	}
	varName := t.UID
	if t.Validation == "" {
		return fmt.Sprintf("!%s", varName)
	}
	return t.Validation
}
