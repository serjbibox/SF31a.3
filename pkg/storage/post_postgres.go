package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/GoNews/pkg/models"
)

type PostPostgres struct {
	db *pgxpool.Pool
}

func newPostPostgres(db *pgxpool.Pool) Post {
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
func (s *PostPostgres) DeletePost(id interface{}) error {
	return nil
}
