package templating

import (
	"fmt"
	"io"
	"strings"
)

// if you're continuing from the read files chapter, you shouldn't redefine this
type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(w io.Writer, p Post) error {
	contentBuilder := strings.Builder{}

	contentBuilder.WriteString(fmt.Sprintf("<h1>%s</h1>\n", p.Title))
	contentBuilder.WriteString(fmt.Sprintf("<p>%s</p>\n", p.Description))
	contentBuilder.WriteString("Tags: <ul>")
	for _, tag := range p.Tags {
		contentBuilder.WriteString(fmt.Sprintf("<li>%s</li>", tag))

	}
	contentBuilder.WriteString("</ul>")

	_, err := fmt.Fprint(w, contentBuilder.String())

	return err
}
