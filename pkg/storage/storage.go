package storage

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/GoNews/pkg/models"
	"github.com/serjbibox/GoNews/pkg/storage/memdb"
	"go.mongodb.org/mongo-driver/mongo"
)

var elog = log.New(os.Stderr, "Storage error\t", log.Ldate|log.Ltime|log.Lshortfile)
var ilog = log.New(os.Stdout, "Storage info\t", log.Ldate|log.Ltime)

//задаёт контракт на работу с таблицей авторов БД.
type Author interface {
	Create(a models.Author) error     // получение всех авторов
	Delete(id string) error           // удаление автора по ID
	GetAll() ([]models.Author, error) // получение всех авторов
}

//задаёт контракт на работу с таблицей публикаций БД.
type Post interface {
	GetAll() ([]models.Post, error)            // получение всех публикаций
	Create(models.Post) (id string, err error) // создание новой публикации
	Update(models.Post) error                  // обновление публикации
	Delete(id string) error                    // удаление публикации по ID
}

// Хранилище данных.
type Storage struct {
	Post
	Author
}

// Конструктор объекта хранилища для БД PostgreSQL.
func NewStoragePostgres(ctx context.Context, db *pgxpool.Pool) (*Storage, error) {
	if ctx == nil {
		elog.Println("context is nil")
		return nil, errors.New("context is nil")
	}
	if db == nil {
		elog.Println("db is nil")
		return nil, errors.New("db is nil")
	}
	return &Storage{
		Post: newPostPostgres(ctx, db),
	}, nil
}

// Конструктор объекта хранилища для БД MemDb.
func NewStorageMemDb(db memdb.DB) (*Storage, error) {
	if db == nil {
		elog.Println("db is nil")
		return nil, errors.New("db is nil")
	}
	return &Storage{
		Post: newPostMemDb(db),
	}, nil
}

// Конструктор объекта хранилища для БД MongoDB.
func NewStorageMongodb(ctx context.Context, db *mongo.Client) (*Storage, error) {
	if ctx == nil {
		elog.Println("context is nil")
		return nil, errors.New("context is nil")
	}
	if db == nil {
		elog.Println("db is nil")
		return nil, errors.New("db is nil")
	}
	return &Storage{
		Post:   newPostMongodb(ctx, db),
		Author: newAuthorMongodb(ctx, db),
	}, nil
}
