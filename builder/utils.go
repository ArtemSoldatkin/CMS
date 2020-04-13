package builder

import (
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
