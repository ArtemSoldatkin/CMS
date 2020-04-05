package tag

import (
	"crypto/rand"
	"fmt"
	"log"
	"strings"
)

type commonError struct {
	message string
}

func (e *commonError) Error() string {
	return e.message
}

func childrenToString(n Tag) string {
	result := ""
	for _, c := range n.Children {
		result += childrenToString(c)
	}
	if n.Name == "input" {
		return fmt.Sprintf("<%s id=\"%s\" %s %s/>", n.Name, n.UID, strings.Join(n.Attributes, " "), n.Value)
	}
	return fmt.Sprintf("<%s id='%s' %s>%s%s</%s>", n.Name, n.UID, strings.Join(n.Attributes, " "), n.Value, result, n.Name)
}

func generateUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("id%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func findChildPos(t *Tag, child Tag) (int, error) {
	for i, c := range t.Children {
		if c.UID == child.UID {
			return i, nil
		}
	}
	return -1, &commonError{"child is not found"}
}
