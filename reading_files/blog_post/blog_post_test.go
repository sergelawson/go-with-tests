package blogpost

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPost(t *testing.T) {
	const (
		post1 = `Title: HTML
Description: How to build a website in HTML
	`
		post2 = `Title: CSS
Description: How to style a website with CSS
	`
	)
	fs := fstest.MapFS{
		"post1.md": {Data: []byte(post1)},
		"post2.md": {Data: []byte(post2)},
	}

	posts, error := NewPostsFromFS(fs)

	if error != nil {
		t.Fatal(error)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	assertPost(t, posts[0], Post{Title: "HTML", Description: "How to build a website in HTML"})
}

func assertPost(t testing.TB, got Post, want Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %+v, got: %+v", want, got)
	}
}
