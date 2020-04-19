package tag

import (
	"crypto/rand"
	"fmt"
	"log"
)

type commonError struct {
	message string
}

func (e *commonError) Error() string {
	return e.message
}

func generateUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("uid%x%x%x%x%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func childrenToString(t Tag) string {
	result := ""
	for _, c := range t.Children {
		result += childrenToString(c)
	}
	if t.Name == "input" {
		return fmt.Sprintf("\n<%s id=\"%s\" %s value=\"%s\"/>", t.Name, t.UID, t.AttributesToString(), t.Value)
	} else if t.Name == "img" {
		return fmt.Sprintf("\n<%s id=\"%s\" %s />", t.Name, t.UID, t.AttributesToString())
	}
	return fmt.Sprintf("\n<%s id=\"%s\" %s>%s%s\n</%s>", t.Name, t.UID, t.AttributesToString(), t.Value, result, t.Name) //t.Value, result, t.Name)
}

// FindChildPosition - find position of child by UID in tag children
func FindChildPosition(t *Tag, child *Tag) (int, error) {
	for i, c := range t.Children {
		if c.UID == child.UID {
			return i, nil
		}
	}
	return -1, &commonError{"Child is not found"}
}
