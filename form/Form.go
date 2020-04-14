package form

import (
	"cms/tag"
)

// CreateForm - create form with input text
func CreateForm(submitLable, url, method string, inputsText []InputText) *tag.Tag {
	submitButton := tag.Tag{Name: "input", Value: submitLable}
	submitButton.Init()
	submitButton.AddAttribute("type", "submit")
	inputs := createTextFieldsToForm(inputsText)
	inputs = append(inputs, submitButton)
	form := tag.Tag{Name: "form", Children: inputs}
	form.Init()
	submitEvent := "e.preventDefault();\n\tconsole.log(\"submited\");"
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

/*


func createFormAction(form *tag.Tag, url, method string) string {
	validation := validateInputs(form)
	values := valuesToObject(form)
	query := makeQueryToServer(queryParams{url, method, values})
	return fmt.Sprintf("\n%s\n%s\n%s\n", preventDefault, validation, query)
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
