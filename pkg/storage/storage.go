package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/GoNews/pkg/models"
	"github.com/serjbibox/GoNews/pkg/storage/memdb"
)

type Post interface {
	Posts() ([]models.Post, error) // получение всех публикаций
	AddPost(models.Post) error     // создание новой публикации
	UpdatePost(models.Post) error  // обновление публикации
	DeletePost(id int) error       // удаление публикации по ID
}

type Storage struct {
	Post
}

func NewStoragePostgres(db *pgxpool.Pool) *Storage {
	return &Storage{
		Post: NewPostPostgres(db),
	}
}

func NewStorageMemDb(db memdb.DB) *Storage {
	return &Storage{
		Post: NewPostMemDb(db),
	}
}
