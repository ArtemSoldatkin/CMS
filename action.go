package main

import "fmt"

type action struct {
	actionType string
	event      string
}

func (a action) toString(tagID string) string {
	return fmt.Sprintf("document.getElementById('%s').addEventListener('%s', function(e){%s})\n", tagID, a.actionType, a.event)
}

func (a *action) addEvent(event string) {
	a.event += fmt.Sprintf("\n%s", event)
}
