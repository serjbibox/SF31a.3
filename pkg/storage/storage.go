package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/GoNews/pkg/models"
	"github.com/serjbibox/GoNews/pkg/storage/memdb"
	"go.mongodb.org/mongo-driver/mongo"
)

type Author interface {
	Create(a models.Author) error
	Delete(id string) error
	GetAll() ([]models.Author, error)
}

type Post interface {
	GetAll() ([]models.Post, error)            // получение всех публикаций
	Create(models.Post) (id string, err error) // создание новой публикации
	Update(models.Post) error                  // обновление публикации
	Delete(id string) error                    // удаление публикации по ID
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
