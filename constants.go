package main

// Actions - dictionary with js actions
var Actions = newActions()

// ActionsType - type of js actions
type ActionsType struct {
	Submit string
	Click  string
}

func newActions() *ActionsType {
	return &ActionsType{
		Submit: "submit",
		Click:  "click",
	}
}
