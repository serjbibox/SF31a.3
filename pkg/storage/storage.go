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
	GetAll() ([]models.Post, error)
	Create(models.Post) (id string, err error)
	Update(models.Post) error
	Delete(id string) error
}

type Storage struct {
	Post
	Author
}

func NewStoragePostgres(ctx context.Context, db *pgxpool.Pool) *Storage {
	return &Storage{
		Post: newPostPostgres(ctx, db),
	}
}

func NewStorageMemDb(db memdb.DB) *Storage {
	return &Storage{
		Post: newPostMemDb(db),
	}
}

func NewStorageMongodb(ctx context.Context, db *mongo.Client) *Storage {
	return &Storage{
		Post:   newPostMongodb(ctx, db),
		Author: newAuthorMongodb(ctx, db),
	}
}
