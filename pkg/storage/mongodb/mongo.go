package mongodb

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DB_USERNAME = "serj"
	DB_HOST     = "192.168.52.129"
	DB_PORT     = "27017"
	DB_NAME     = "mongodb"
	DB_SSLMODE  = "require"
)

var elog = log.New(os.Stderr, "mongodb error\t", log.Ldate|log.Ltime|log.Lshortfile)
var ilog = log.New(os.Stdout, "mongodb info\t", log.Ldate|log.Ltime)

func New(ctx context.Context) (*mongo.Client, error) {
	mongoOpts := options.Client().ApplyURI(DB_NAME + "://" + DB_HOST + ":" + DB_PORT + "/")
	client, err := mongo.Connect(ctx, mongoOpts)
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	ilog.Println("connected to mongo database")
	return client, err
}
