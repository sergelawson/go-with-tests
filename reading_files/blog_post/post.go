package blogpost

import (
	"bufio"
	"io"
	"strings"
)

type Post struct {
	Title       string
	Description string
	/* , Description, Body string
	Tags                     []string */
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readLine(titleSeparator)
	description := readLine(descriptionSeparator)

	post := Post{Title: title, Description: description}

	return post, nil
}
