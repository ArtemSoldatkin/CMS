package form

import (
	"cms/tag"
	"fmt"
)

// InputText - text field for form
type InputText struct {
	Label, Placeholder, ValueName, Validation string
}

func createTextField(params InputText) []tag.Tag {
	input := tag.Tag{Name: "input", Value: "", ValueName: params.ValueName}
	input.Init()
	input.AddAttribute("type", "text")
	if params.Placeholder != "" {
		input.AddAttribute("placeholder", params.Placeholder)
	}
	event := fmt.Sprintf("%s=e.target.value", tag.UIDToValueName(input.UID))
	input.AddAction("change", event)
	if params.Label != "" {
		label := tag.Tag{Name: "label", Value: params.Label}
		label.Init()
		label.AddAttribute("for", input.UID)
		return []tag.Tag{label, input}
	}
	return []tag.Tag{input}
}
