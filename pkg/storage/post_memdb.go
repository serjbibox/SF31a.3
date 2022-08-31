package storage

import (
	"errors"
	"strconv"

	"github.com/serjbibox/GoNews/pkg/models"
	"github.com/serjbibox/GoNews/pkg/storage/memdb"
)

//Объект, реализующий интерфейс работы с публикациями в памяти.
type PostMemdb struct {
	db memdb.DB
}

//Конструктор PostMemdb
func newPostMemDb(db memdb.DB) Post {
	return &PostMemdb{db: db}
}

// получение всех публикаций
func (s *PostMemdb) GetAll() ([]models.Post, error) {
	return s.db, nil
}

// создание новой публикации
func (s *PostMemdb) Create(p models.Post) (string, error) {
	id, err := strconv.Atoi(s.db[len(s.db)-1].ID)
	if err != nil {
		return "", err
	}
	p.ID = strconv.Itoa(id + 1)
	s.db = append(s.db, p)
	return s.db[len(s.db)-1].ID, nil
}

// обновление публикации
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

// удаление публикации по ID
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
