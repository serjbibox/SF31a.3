package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/GoNews/pkg/models"
)

const (
	ID_TYPE = 1 + iota
	NAME_TYPE
)

type PostPostgres struct {
	db *pgxpool.Pool
}

func NewPostPostgres(db *pgxpool.Pool) Post {
	return &PostPostgres{db: db}
}

func (s *PostPostgres) Posts() ([]models.Post, error) {
	return posts, nil
}

func (s *PostPostgres) AddPost(models.Post) error {
	return nil
}
func (s *PostPostgres) UpdatePost(models.Post) error {
	return nil
}
func (s *PostPostgres) DeletePost(id int) error {
	return nil
}
