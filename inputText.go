package main

import (
	"fmt"
	"strings"
)

type inputText struct {
	label, placeholder, valueName, validation string
}

func createTextField(params inputText) []tag {
	valueName := params.valueName
	if valueName == "" {
		valueName = params.label
	}
	input := tag{name: "input", attributes: []string{"type=\"text\""}, valueName: valueName}
	input.Init()
	if params.placeholder != "" {
		input.addAttribute(fmt.Sprintf("placeholder=\"%s\"", params.placeholder))
	}
	if params.label != "" {
		label := tag{name: "label", value: params.label, attributes: []string{fmt.Sprintf("for=%s", input.uid)}}
		label.Init()
		return []tag{label, input}
	}
	return []tag{input}
}

func onChange(input *tag) string {
	return fmt.Sprintf("document.getElementById(\"%s\").addEventListener(\"change\",function(e){%s=e.target.value})\n", input.uid, strings.ReplaceAll(input.uid, "-", ""))
}

func checkTextInput(t *tag) bool {
	return t.name == "input" && checkAttribute(t, "text")
}

func createValidation(t *tag) string {
	if t.validation == "" {
		return fmt.Sprintf("!%s", getVariable(t))
	}
	return t.validation
}
