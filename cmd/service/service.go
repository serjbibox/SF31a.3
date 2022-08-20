package main

import (
	"log"
	"os"

	"github.com/serjbibox/GoNews/pkg/storage"
	"github.com/serjbibox/GoNews/pkg/storage/postgresql"
)

var elog = log.New(os.Stderr, "service error\t", log.Ldate|log.Ltime|log.Lshortfile)
var ilog = log.New(os.Stdout, "service info\t", log.Ldate|log.Ltime)

func main() {
	db, err := postgresql.New(postgresql.GetConnectionString())
	if err != nil {

	}
	s := storage.NewStoragePostgres(db)
	s.Posts()
}
