package blogpost

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPost(t *testing.T) {
	fs := fstest.MapFS{
		"hello.md": {Data: []byte("Title: Post 1")},
		"world.md": {Data: []byte("Title: World")},
	}

	posts, error := NewPostsFromFS(fs)

	got := posts[0]

	want := Post{Title: "Post 1"}

	if error != nil {
		t.Fatal(error)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %+v, got: %+v", want, got)
	}
}
