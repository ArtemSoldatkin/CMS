package form

import (
	"cms/tag"
	"fmt"
)

// CreateForm - create form with input text
func CreateForm(submitLable, url string, inputsText []InputText) *tag.Tag {
	submitButton := tag.Tag{Name: "input", Value: submitLable}
	submitButton.Init()
	submitButton.AddAttribute("type", "submit")
	inputs := createTextFieldsToForm(inputsText)
	inputs = append(inputs, submitButton)
	form := tag.Tag{Name: "form", Children: inputs}
	form.Init()
	submitEvent := fmt.Sprintf("e.preventDefault();\n%s\n%s", createFormValidation(&form), createRequest(&form, url))
	form.AddAction("submit", submitEvent)
	return &form
}

func createTextFieldsToForm(inputs []InputText) []tag.Tag {
	var result []tag.Tag
	for _, input := range inputs {
		result = append(result, createTextField(input)...)
	}
	return result
}
