package main

import (
	"cms/tag"
	"fmt"
	"strings"
)

const preventDefault = "e.preventDefault()"

type queryParams struct {
	url, method, data string
}

func makeFetch(qp *queryParams) string {
	return fmt.Sprintf("fetch('%s',{method: '%s',headers: {'Content-Type': 'application/json',},body: JSON.stringify(%s)})", qp.url, qp.method, qp.data)
}

func returnFetchResult() string {
	return ".then((response) => response.json())\n.then((data)=>{console.log('Success:', data);})\n.catch((error) => {console.error('Error:', error);});"
}

func makeQueryToServer(qp queryParams) string {
	return fmt.Sprintf("%s\n%s", makeFetch(&qp), returnFetchResult())
}

func valuesToObject(form *tag.Tag) string {
	var variables, variableNames []string
	for _, c := range form.Children {
		if checkTextInput(&c) {
			variables = append(variables, getVariable(&c))
			variableNames = append(variableNames, c.ValueName)
		}
	}
	var result []string
	for i, v := range variables {
		if variableNames[i] != "" {
			result = append(result, fmt.Sprintf("%s: %s", variableNames[i], v))
		} else {
			continue
			//result = append(result, fmt.Sprintf("${%s}", v))
		}
	}
	return fmt.Sprintf("{\n%s\n}", strings.Join(result, ",\n"))
}
