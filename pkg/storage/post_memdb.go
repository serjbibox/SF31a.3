package storage

import (
	"errors"
	"strconv"

	"github.com/serjbibox/GoNews/pkg/models"
	"github.com/serjbibox/GoNews/pkg/storage/memdb"
)

type PostMemdb struct {
	db memdb.DB
}

func newPostMemDb(db memdb.DB) Post {
	return &PostMemdb{db: db}
}

func (s *PostMemdb) GetAll() ([]models.Post, error) {
	return s.db, nil
}

func (s *PostMemdb) Create(p models.Post) (string, error) {
	id, err := strconv.Atoi(s.db[len(s.db)-1].ID)
	if err != nil {
		return "", err
	}
	p.ID = strconv.Itoa(id + 1)
	s.db = append(s.db, p)
	return s.db[len(s.db)-1].ID, nil
}
func (s *PostMemdb) Update(p models.Post) error {
	id, err := strconv.Atoi(p.ID)
	if err != nil {
		return err
	}
	if id >= len(s.db) || id == 0 {
		return errors.New("wrong post id")
	}
	s.db[id-1] = p
	return nil
}
func (s *PostMemdb) Delete(id string) error {
	delId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	if delId >= len(s.db) || delId == 0 {
		return errors.New("wrong post id")
	}
	s.db[delId-1] = models.Post{}
	return nil
}
