package form

/*
import (
	"cms/actions"
	"cms/tag"
	"fmt"
	"strings"
)

// CreateForm - create form with input text
func CreateForm(submitLable, url, method string, inputsText []InputText) *tag.Tag {
	submitButton := tag.Tag{Name: "input", Value: submitLable, Attributes: []string{"type=\"submit\""}}
	submitButton.Init()
	inputs := createTextFields(inputsText)
	inputs = append(inputs, submitButton)
	form := tag.Tag{Name: "form", Children: inputs}
	form.Init()
	submitAction := actions.Actions{"submit", createFormAction(&form, url, method)}
	form.AddAction(submitAction)
	return &form
}

func createFormAction(form *tag.Tag, url, method string) string {
	validation := validateInputs(form)
	values := valuesToObject(form)
	query := makeQueryToServer(queryParams{url, method, values})
	return fmt.Sprintf("\n%s\n%s\n%s\n", preventDefault, validation, query)
}

func createTextFields(inputs []InputText) []tag.Tag {
	var result []tag.Tag
	for _, input := range inputs {
		result = append(result, createTextField(input)...)
	}
	return result
}

func validateInputs(f *tag.Tag) string {
	var values []string
	for _, c := range f.Children {
		if checkTextInput(&c) {
			values = append(values, createValidation(&c))
		}
	}
	condition := strings.Join(values, "||")
	return fmt.Sprintf("if(%s)alert(\"some fields is empty\")", condition)
}
*/
