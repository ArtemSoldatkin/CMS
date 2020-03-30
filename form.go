package main

import (
	"fmt"
	"strings"
)

func createForm(submitLable string, inputsText []inputText) *tag {
	submitButton := tag{name: "input", value: submitLable, attributes: []string{"type=\"submit\""}}
	submitButton.Init()
	inputs := createTextFields(inputsText)
	inputs = append(inputs, submitButton)
	form := tag{name: "form", children: inputs}
	form.Init()
	submitAction := action{"submit", createFormAction(&form)}
	form.addAction(submitAction)
	return &form
}

func createFormAction(form *tag) string {
	var variables []string
	for _, c := range form.children {
		if checkTextInput(&c) {
			variables = append(variables, getVariable(&c))
		}
	}
	return fmt.Sprintf("alert([%s].join(\" \"))", strings.Join(variables, ","))
}

func createTextFields(inputs []inputText) []tag {
	var result []tag
	for _, input := range inputs {
		result = append(result, createTextField(input)...)
	}
	return result
}
