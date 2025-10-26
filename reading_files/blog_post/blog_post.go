package blogpost

import (
	"io/fs"
)

func NewPostsFromFS(filesystem fs.FS) ([]Post, error) {
	posts := []Post{}

	dir, err := fs.ReadDir(filesystem, ".")

	if err != nil {
		return nil, err
	}

	for _, file := range dir {
		post, err := getPost(filesystem, file.Name())

		if err != nil {

			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, err
}

func getPost(filesystem fs.FS, fileName string) (Post, error) {

	postFile, err := filesystem.Open(fileName)

	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	return newPost(postFile)
}
