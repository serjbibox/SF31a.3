package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/GoNews/pkg/models"
	"github.com/serjbibox/GoNews/pkg/storage/memdb"
	"go.mongodb.org/mongo-driver/mongo"
)

type Author interface {
	AddAuthor(a models.Author) error
	DeleteAuthor(id interface{}) error
	Authors() ([]models.Author, error)
}

type Post interface {
	Posts() ([]models.Post, error)   // получение всех публикаций
	AddPost(models.Post) error       // создание новой публикации
	UpdatePost(models.Post) error    // обновление публикации
	DeletePost(id interface{}) error // удаление публикации по ID
}

type Storage struct {
	Post
	Author
}

func NewStoragePostgres(db *pgxpool.Pool) *Storage {
	return &Storage{
		Post: newPostPostgres(db),
	}
}

func NewStorageMemDb(db memdb.DB) *Storage {
	return &Storage{
		Post: newPostMemDb(db),
	}
}

func NewStorageMongodb(db *mongo.Client, ctx context.Context) *Storage {
	return &Storage{
		Post:   newPostMongodb(db, ctx),
		Author: newAuthorMongodb(db, ctx),
	}
}
