package main

import (
	"fmt"
	"io/ioutil"
)

const buildPath string = "build"

func writeToFile(fileName string, data string) {
	b := []byte(data)
	err := ioutil.WriteFile(fmt.Sprintf("%s/%s.html", buildPath, fileName), b, 0644)
	if err != nil {
		panic(err)
	}
}
