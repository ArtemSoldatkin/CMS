package main

import "fmt"

type action struct {
	actionType string
	event      string
}

func (a action) toString(tagID string) string {
	return fmt.Sprintf("document.getElementById('%s').addEventListener('%s', function(){%s})\n", tagID, a.actionType, a.event)
}
