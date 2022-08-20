package main

import (
	"context"
	"log"
	"os"

	"github.com/serjbibox/GoNews/pkg/storage"
	"github.com/serjbibox/GoNews/pkg/storage/mongodb"
)

var elog = log.New(os.Stderr, "service error\t", log.Ldate|log.Ltime|log.Lshortfile)
var ilog = log.New(os.Stdout, "service info\t", log.Ldate|log.Ltime)
var ctx = context.Background()

func main() {
	db, err := mongodb.New(ctx)
	if err != nil {
		elog.Println(err)
	}
	defer db.Disconnect(ctx)
	s := storage.NewStorageMongodb(db, ctx)
	/*databases, err := db.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	ilog.Println("dabases: ", databases)
	log.Println("insert", s.AddPost(models.Post{
		AuthorName: "Petro",
		Content:    "new post",
	}))*/
	ilog.Println(s.Posts())
	//s := storage.NewStoragePostgres(db)
	//s.Posts()
}
