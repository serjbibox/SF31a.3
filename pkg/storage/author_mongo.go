package storage

import (
	"context"
	"errors"
	"time"

	"github.com/serjbibox/GoNews/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthorMongodb struct {
	db  *mongo.Client
	ctx context.Context
}

func newAuthorMongodb(db *mongo.Client, ctx context.Context) Author {
	return &AuthorMongodb{
		db:  db,
		ctx: ctx,
	}
}

func (s *AuthorMongodb) Authors() ([]models.Author, error) {
	collection := s.db.Database(MONGO_NEWS_DB).Collection(MONGO_AUTHORS)
	filter := bson.D{}
	cur, err := collection.Find(s.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(s.ctx)
	var data []models.Author
	for cur.Next(s.ctx) {
		var l models.Author
		err := cur.Decode(&l)
		if err != nil {
			return nil, err
		}
		data = append(data, l)
	}
	return data, cur.Err()
}

func (s *AuthorMongodb) AddAuthor(a models.Author) error {
	collection := s.db.Database(MONGO_NEWS_DB).Collection(MONGO_AUTHORS)
	a.MongoID = primitive.NewObjectIDFromTimestamp(time.Now()).Hex()
	_, err := collection.InsertOne(s.ctx, a)
	if err != nil {
		return err
	}
	return nil
}

func (s *AuthorMongodb) DeleteAuthor(id interface{}) error {
	id, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	collection := s.db.Database(MONGO_NEWS_DB).Collection(MONGO_AUTHORS)
	res, err := collection.DeleteOne(s.ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no tasks to delete")
	}
	return nil
}
