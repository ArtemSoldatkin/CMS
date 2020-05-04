package tag

import (
	"fmt"
	"log"
	"strings"
)

// Tag - html Tag struct
type Tag struct {
	UID        string
	Name       string            `json:"name"`
	Attributes map[string]string `json:"attributes"`
	Style      map[string]string `json:"style"`
	Children   []*Tag            `json:"children"`
	Value      string            `json:"value"`
	ValueName  string            `json:"value-name"`
	Validation string            `json:"validation"`
	Actions    map[string]string
}

// Init - initialize tag
func (t *Tag) Init() {
	t.UID = generateUID()
	t.Attributes = make(map[string]string)
	t.Style = make(map[string]string)
	t.Actions = make(map[string]string)
}

func (t *Tag) String() string {
	return childrenToString(t)
}

// Attributes

// AddAttribute - add attribute to tag
func (t *Tag) AddAttribute(key, value string) {
	t.Attributes[key] = value
}

// RemoveAttribute - remove attribute from tag
func (t *Tag) RemoveAttribute(key string) {
	delete(t.Attributes, key)
}

// AttributesToString - convert attributes to string
func (t Tag) AttributesToString() string {
	var result string
	for k, v := range t.Attributes {
		result += fmt.Sprintf("%s=\"%s\" ", k, v)
	}
	return strings.TrimRight(result, " ")
}

// Style

// AddStyle - add style to tag
func (t *Tag) AddStyle(key, value string) {
	t.Style[key] = value
}

// RemoveStyle - remove style from tag
func (t *Tag) RemoveStyle(key string) {
	delete(t.Style, key)
}

// StyleToString - convert style to string
func (t Tag) StyleToString() string {
	var result string
	for k, v := range t.Style {
		result += fmt.Sprintf("\t%s: %s;\n", k, v)
	}
	if result != "" {
		return fmt.Sprintf("#%s {\n%s}", t.UID, result)
	}
	return result
}

// Children

// AppendChild - add new child to tag
func (t *Tag) AppendChild(child *Tag) {
	t.Children = append(t.Children, child)
}

// RemoveChild - remove child from tag children
func (t *Tag) RemoveChild(UID string, child *Tag) {
	pos, err := FindChildPosition(t, child)
	if err != nil {
		log.Fatal(err)
	}
	copy(t.Children[pos:], t.Children[pos+1:])
	t.Children[len(t.Children)-1] = nil
	t.Children = t.Children[:len(t.Children)-1]
}

// SwitchChildren - switch child position in tag childnren
func (t *Tag) SwitchChildren(position int, child *Tag) {
	currPos, err := FindChildPosition(t, child)
	if err != nil {
		log.Fatal(err)
	}
	t.Children = append(t.Children[:currPos], t.Children[currPos+1:]...)
	t.Children = append(t.Children, nil)
	copy(t.Children[position+1:], t.Children[position:])
	t.Children[position] = child
}

// Actions

// AddAction - add action to tag
func (t *Tag) AddAction(key, value string) {
	t.Actions[key] = value
}

// RemoveAction - remove action from tag
func (t *Tag) RemoveAction(key string) {
	delete(t.Actions, key)
}

// ActionToString - convert actions to string
func (t Tag) ActionToString() string {
	var result string
	for k, v := range t.Actions {
		result += fmt.Sprintf("%s\n", CreateEventListener(t.UID, k, v))
	}
	return strings.TrimRight(result, "\n")
}

// CreateValidation - create validation to tag value
func (t Tag) CreateValidation() string {
	if t.Validation == "" {
		return fmt.Sprintf("!%s", t.UID)
	}
	return fmt.Sprintf("%s !== %s", t.UID, t.Validation)
}
