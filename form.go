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
	var variables, variableNames []string
	for _, c := range form.children {
		if checkTextInput(&c) {
			variables = append(variables, getVariable(&c))
			variableNames = append(variableNames, c.valueName)
		}
	}
	validation := validateInputs(form)
	var result []string
	for i, v := range variables {
		if variableNames[i] != "" {
			result = append(result, fmt.Sprintf("%s: ${%s}", variableNames[i], v))
		} else {
			result = append(result, fmt.Sprintf("${%s}", v))
		}
	}

	query := makeQueryToServer(queryParams{"http://localhost:5000", "POST", "data"})

	return fmt.Sprintf("\n%s\n%s\n%s\nalert(`%s`)", preventDefault, validation, query, strings.Join(result, ", "))
}

func createTextFields(inputs []inputText) []tag {
	var result []tag
	for _, input := range inputs {
		result = append(result, createTextField(input)...)
	}
	return result
}

func validateInputs(f *tag) string {
	var values []string
	for _, c := range f.children {
		if checkTextInput(&c) {
			values = append(values, createValidation(&c))
		}
	}
	condition := strings.Join(values, "||")
	return fmt.Sprintf("if(%s)alert(\"some fields is empty\")", condition)
}
