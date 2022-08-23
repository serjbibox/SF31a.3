package storage

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/GoNews/pkg/models"
)

type PostPostgres struct {
	db *pgxpool.Pool
}

func newPostPostgres(db *pgxpool.Pool) Post {
	return &PostPostgres{db: db}
}

func (s *PostPostgres) GetAll() ([]models.Post, error) {
	return posts, nil
}

func (s *PostPostgres) Create(p models.Post) (string, error) {
	var id int
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO posts (title, content, author_id)
		VALUES ($1, $2, $3) RETURNING id;
		`,
		p.Title,
		p.Content,
		p.AuthorID,
	).Scan(&id)
	return strconv.Itoa(id), err
}
func (s *PostPostgres) Update(models.Post) error {
	return nil
}
func (s *PostPostgres) Delete(id string) error {
	return nil
}
