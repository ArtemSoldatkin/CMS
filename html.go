package main

import "fmt"

type html struct {
	children []tag
}

func (h html) String() string {
	result := ""
	for _, c := range h.children {
		result += childrenToString(c)
	}
	return fmt.Sprintf("<!DOCTYPE html><html><head></head><body>%s</body></html>", result)
}
