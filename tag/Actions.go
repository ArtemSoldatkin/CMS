package tag

import "fmt"

/*
// Action - Event in js.
type Action struct {
	Type, Event string
}
*/

/*
// ToString ...
func (a Actions) ToString(tagID string) string {
	return fmt.Sprintf("document.getElementById('%s').addEventListener('%s', function(e){%s})\n", tagID, a.Type, a.Event)
}

// AddEvent ...
func (a *Actions) AddEvent(event string) {
	a.Event += fmt.Sprintf("\n%s", event)
}
*/

// CreateEventListener - function to create eventListener.
func CreateEventListener(uid, actionType, event string) string {
	return fmt.Sprintf("document.getElementById(\"%s\").addEventListener(\"%s\", function(e){\n\t%s\n})", uid, actionType, event)
}
