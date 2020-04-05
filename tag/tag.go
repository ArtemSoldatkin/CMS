package tag

import actions "cms/actions"

// Tag - html Tag struct
type Tag struct {
	UID        string
	Name       string
	Attributes []string
	Style      []string
	Children   []Tag
	Value      string
	ValueName  string
	Actions    []actions.Actions
	Validation string
}

// Init - initialize tag
func (t *Tag) Init() {
	t.UID = generateUID()
}

func (t Tag) String() string {
	return childrenToString(t)
}

// AddAttribute - add attribute to tag
func (t *Tag) AddAttribute(attr string) {
	t.Attributes = append(t.Attributes, attr)
}

// AddAction - add action to tag
func (t *Tag) AddAction(act actions.Actions) {
	t.Actions = append(t.Actions, act)
}

// AppendChildren ...
func (t *Tag) AppendChildren(newChildren []Tag) {
	t.Children = append(t.Children, newChildren...)
}

// RemoveChildren ...
func (t *Tag) RemoveChildren(childrenToRemove []Tag) {
	var result []Tag
	for _, ctr := range childrenToRemove {
		for _, c := range t.Children {
			if ctr.UID != c.UID {
				result = append(result, c)
			}
		}
	}
	t.Children = result
}

// SwitchChildren ...
func (t *Tag) SwitchChildren(child Tag, pos int) {
	currPos, ok := findChildPos(t, child)
	if ok != nil {
		return
	}
	t.Children = append(t.Children[:currPos], t.Children[currPos+1:]...)
	t.Children = append(t.Children, Tag{})
	copy(t.Children[pos+1:], t.Children[pos:])
	t.Children[pos] = child
}
