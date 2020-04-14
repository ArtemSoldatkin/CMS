package builder

import (
	"cms/tag"
	"fmt"
	"strings"
)

func createHead(title string, args ...string) string {
	return fmt.Sprintf("<head>\n<title>%s</title>\n%s\n</head>", title, strings.Join(args, "\n"))
}

func createCSSLink(path string) string {
	return fmt.Sprintf("<link rel='stylesheet' href='%s.css'>", path)
}

func createScriptLink(path string) string {
	return fmt.Sprintf("<script type='text/javascript' src='%s.js'></script>", path)
}

func createBody(children *[]string) string {
	return fmt.Sprintf("<body>\n%s\n</body>", strings.Join(*children, "\n"))
}

func createVariables(t *tag.Tag) string {
	var result string
	for _, c := range t.Children {
		result += createVariables(&c)
	}
	if t.Name != "input" {
		return result
	}
	return fmt.Sprintf("%s, %s", result, tag.UIDToValueName(t.UID))
}

func createEventListeners(t *tag.Tag) string {
	var result string
	for _, c := range t.Children {
		result += createEventListeners(&c)
	}
	return fmt.Sprintf("%s\n%s", result, t.ActionToString())
}
