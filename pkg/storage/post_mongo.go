package storage

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/serjbibox/GoNews/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGO_NEWS_DB = "gonews"
	MONGO_POSTS   = "posts"
	MONGO_AUTHORS = "authors"
)

type PostMongodb struct {
	db  *mongo.Client
	ctx context.Context
}

func newPostMongodb(ctx context.Context, db *mongo.Client) Post {
	return &PostMongodb{
		db:  db,
		ctx: ctx,
	}
}

func (s *PostMongodb) GetAll() ([]models.Post, error) {
	collection := s.db.Database(MONGO_NEWS_DB).Collection(MONGO_POSTS)
	filter := bson.D{}
	cur, err := collection.Find(s.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(s.ctx)
	var data []models.Post
	for cur.Next(s.ctx) {
		var p models.Post
		err := cur.Decode(&p)
		if err != nil {
			return nil, err
		}

		collection := s.db.Database(MONGO_NEWS_DB).Collection(MONGO_AUTHORS)
		filter := bson.D{primitive.E{Key: "_id", Value: p.AuthorID}}
		cur := collection.FindOne(s.ctx, filter)
		var a models.Author
		err = cur.Decode(&a)
		if err != nil {
			return nil, err
		}
		p.AuthorName = a.Name

		data = append(data, p)
	}
	return data, cur.Err()
}

func (s *PostMongodb) Create(p models.Post) (string, error) {
	collection := s.db.Database(MONGO_NEWS_DB).Collection(MONGO_AUTHORS)
	filter := bson.D{primitive.E{Key: "_id", Value: p.AuthorID}}
	cur := collection.FindOne(s.ctx, filter)
	var a models.Author
	err := cur.Decode(&a)
	if err != nil {
		return "", err
	}

	p.AuthorName = a.Name
	p.CreatedAt = int64(primitive.Timestamp{T: uint32(time.Now().Unix())}.T)

	collection = s.db.Database(MONGO_NEWS_DB).Collection(MONGO_POSTS)
	p.ID = primitive.NewObjectIDFromTimestamp(time.Now()).Hex()
	_, err = collection.InsertOne(s.ctx, p)
	if err != nil {
		return "", err
	}
	log.Println("add new post with id:", p.ID)
	return p.ID, nil
}

func (s *PostMongodb) Update(p models.Post) error {
	collection := s.db.Database(MONGO_NEWS_DB).Collection(MONGO_POSTS)
	res, err := collection.UpdateOne(
		s.ctx,
		bson.M{"_id": p.ID},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "title", Value: p.Title},
				{Key: "content", Value: p.Content},
				{Key: "author_id", Value: p.AuthorID},
			}},
		},
	)
	if err != nil {
		return err
	}
	log.Println(res)
	return nil
}

func (s *PostMongodb) Delete(id string) error {
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	collection := s.db.Database(MONGO_NEWS_DB).Collection(MONGO_POSTS)
	res, err := collection.DeleteOne(s.ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no post to delete")
	}
	return nil
}
