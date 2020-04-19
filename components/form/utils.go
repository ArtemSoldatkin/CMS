package form

import (
	"cms/tag"
	"fmt"
	"strings"
)

func getVariables(t *tag.Tag, result *[]string) {
	for _, c := range t.Children {
		getVariables(&c, result)
	}
	if t.Name == "input" && t.ValueName != "" {
		*result = append(*result, t.UID)
	}
}

func getValidations(t *tag.Tag) string {
	var result string
	for _, c := range t.Children {
		result += getValidations(&c)
	}
	if t.Name != "input" || t.ValueName == "" {
		return result
	}
	return fmt.Sprintf("%s || %s", result, t.CreateValidation())
}

func createFormValidation(t *tag.Tag) string {
	validation := getValidations(t)
	if validation == "" {
		return ""
	}
	validation = strings.TrimLeft(validation, " || ")
	falseEvent := "\nreturn;"
	return fmt.Sprintf("if(%s){\n%s\n}", validation, falseEvent)
}

func getInputValues(t *tag.Tag) string {
	var result string
	for _, c := range t.Children {
		result += getInputValues(&c)
	}
	if t.Name != "input" || t.ValueName == "" {
		return result
	}
	return fmt.Sprintf("%s,\n%s: %s", result, t.ValueName, t.UID)
}

func getDataToRequest(t *tag.Tag) string {
	if t.Name != "form" {
		return ""
	}
	return fmt.Sprintf("{\n%s\n}", strings.TrimLeft(getInputValues(t), ","))
}

func createQuery(url, data string) string {
	return fmt.Sprintf("fetch(\"%s\",{method: \"POST\",headers: {'Content-Type': 'application/json',},body: JSON.stringify(%s)})", url, data)
}

func createResponse() string {
	return ".then((response) => response.json())\n.then((data)=>{console.log('Success:', data);})\n.catch((error) => {console.error('Error:', error);});"
}

func createRequest(t *tag.Tag, url string) string {
	data := getDataToRequest(t)
	return fmt.Sprintf("%s\n%s", createQuery(url, data), createResponse())
}
