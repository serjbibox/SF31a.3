package storage

import (
	"context"
	"errors"
	"fmt"
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

func newPostMongodb(db *mongo.Client, ctx context.Context) Post {
	return &PostMongodb{
		db:  db,
		ctx: ctx,
	}
}

func (s *PostMongodb) Posts() ([]models.Post, error) {
	collection := s.db.Database(MONGO_NEWS_DB).Collection(MONGO_POSTS)
	filter := bson.D{}
	cur, err := collection.Find(s.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(s.ctx)
	var data []models.Post
	for cur.Next(s.ctx) {
		var l models.Post
		err := cur.Decode(&l)
		if err != nil {
			return nil, err
		}
		data = append(data, l)
	}
	return data, cur.Err()
}

func (s *PostMongodb) AddPost(p models.Post) error {
	collection := s.db.Database(MONGO_NEWS_DB).Collection(MONGO_POSTS)
	p.MongoID = primitive.NewObjectIDFromTimestamp(time.Now()).Hex()
	_, err := collection.InsertOne(s.ctx, p)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostMongodb) UpdatePost(p models.Post) error {
	collection := s.db.Database(MONGO_NEWS_DB).Collection(MONGO_POSTS)
	id, _ := primitive.ObjectIDFromHex(p.MongoID)
	result, err := collection.UpdateOne(
		s.ctx,
		bson.M{"_id": id},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "title", Value: p.Title},
				{Key: "content", Value: p.Content},
				{Key: "author_id", Value: p.AuthorID},
				{Key: "name", Value: p.AuthorName},
			}},
		},
	)
	if err != nil {
		return err
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)

	return nil
}

func (s *PostMongodb) DeletePost(id interface{}) error {
	id, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return err
	}
	filter := bson.D{primitive.E{Key: "_id", Value: id}}
	collection := s.db.Database(MONGO_NEWS_DB).Collection(MONGO_POSTS)
	res, err := collection.DeleteOne(s.ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.New("no tasks to delete")
	}
	return nil
}
