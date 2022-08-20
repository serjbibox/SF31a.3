package storage

import (
	"context"
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

func newPostMongodb(db *mongo.Client, ctx context.Context) Post {
	return &PostMongodb{
		db:  db,
		ctx: ctx,
	}
}

func (s *PostMongodb) Posts() ([]models.Post, error) {
	collection := s.db.Database(MONGO_NEWS_DB).Collection(MONGO_POSTS)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var data []models.Post
	for cur.Next(context.Background()) {
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
	p.MID = primitive.NewObjectIDFromTimestamp(time.Now())
	res, err := collection.InsertOne(context.Background(), p)
	if err != nil {
		return err
	}
	log.Println(res.InsertedID.(primitive.ObjectID))
	return nil
}

func (s *PostMongodb) UpdatePost(models.Post) error {
	return nil
}

func (s *PostMongodb) DeletePost(id int) error {
	return nil
}
