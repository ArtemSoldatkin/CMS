package main

import "fmt"

const preventDefault = "e.preventDefault()"

type queryParams struct {
	url, method, data string
}

func makeFetch(qp *queryParams) string {
	return fmt.Sprintf("fetch('%s',{method: '%s',headers: {'Content-Type': 'application/json',},body: JSON.stringify(\"%s\")})", qp.url, qp.method, qp.data)
}

func returnFetchResult() string {
	return ".then((response) => response.json())\n.then((data)=>{console.log('Success:', data);})\n.catch((error) => {console.error('Error:', error);});"
}

func makeQueryToServer(qp queryParams) string {
	return fmt.Sprintf("%s\n%s", makeFetch(&qp), returnFetchResult())
}
