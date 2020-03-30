package main

type tag struct {
	uid        string
	name       string
	attributes []string
	style      []string
	children   []tag
	value      string
	actions    []action
}

func (t *tag) Init() {
	t.uid = generateUID()
}

func (t tag) String() string {
	return childrenToString(t)
}

func (t *tag) addAttribute(attr string) {
	t.attributes = append(t.attributes, attr)
}

func (t *tag) addAction(act action) {
	t.actions = append(t.actions, act)
}

func (t *tag) appendChildren(newChildren []tag) {
	t.children = append(t.children, newChildren...)
}

func (t *tag) removeChildren(childrenToRemove []tag) {
	var result []tag
	for _, ctr := range childrenToRemove {
		for _, c := range t.children {
			if ctr.uid != c.uid {
				result = append(result, c)
			}
		}
	}
	t.children = result
}

func (t *tag) switchChildren(child tag, pos int) {
	currPos, ok := findChildPos(t, child)
	if ok != nil {
		return
	}
	t.children = append(t.children[:currPos], t.children[currPos+1:]...)
	t.children = append(t.children, tag{})
	copy(t.children[pos+1:], t.children[pos:])
	t.children[pos] = child
}
