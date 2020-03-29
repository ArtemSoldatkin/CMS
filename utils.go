package main

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

func childrenToString(n tag) string {
	result := ""
	for _, c := range n.children {
		result += childrenToString(c)
	}
	return fmt.Sprintf("<%s id='%s' %s>%s%s</%s>", n.name, n.uid, strings.Join(n.attributes, " "), n.value, result, n.name)
}

func generateUID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func findChildPos(t *tag, child tag) (int, error) {
	for i, c := range t.children {
		if c.uid == child.uid {
			return i, nil
		}
	}
	return -1, &commonError{"child is not found"}
}
