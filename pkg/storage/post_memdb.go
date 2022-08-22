package storage

import (
	"github.com/serjbibox/GoNews/pkg/models"
	"github.com/serjbibox/GoNews/pkg/storage/memdb"
)

type PostMemdb struct {
	db memdb.DB
}

func newPostMemDb(db memdb.DB) Post {
	return &PostMemdb{db: db}
}

func (s *PostMemdb) Posts() ([]models.Post, error) {
	return posts, nil
}

func (s *PostMemdb) AddPost(models.Post) error {
	return nil
}
func (s *PostMemdb) UpdatePost(models.Post) error {
	return nil
}
func (s *PostMemdb) DeletePost(id interface{}) error {
	return nil
}

var posts = []models.Post{
	{
		//ID:      1,
		Title:   "Effective Go",
		Content: "Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand.",
	},
	{
		//ID:      2,
		Title:   "The Go Memory Model",
		Content: "The Go memory model specifies the conditions under which reads of a variable in one goroutine can be guaranteed to observe values produced by writes to the same variable in a different goroutine.",
	},
}
