package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/serjbibox/GoNews/pkg/handler"
	"github.com/serjbibox/GoNews/pkg/storage"
	"github.com/serjbibox/GoNews/pkg/storage/mongodb"
)

var elog = log.New(os.Stderr, "service error\t", log.Ldate|log.Ltime|log.Lshortfile)
var ilog = log.New(os.Stdout, "service info\t", log.Ldate|log.Ltime)
var ctx = context.Background()

const (
	HTTP_PORT = "8080"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func main() {
	var err error

	db, err := mongodb.New(ctx)
	if err != nil {
		elog.Println(err)
	}
	defer db.Disconnect(ctx)
	s := storage.NewStorageMongodb(ctx, db)

	//db, err := postgresql.New(postgresql.GetConnectionString())
	//if err != nil {
	//	elog.Println(err)
	//}
	//defer db.Close()
	//s := storage.NewStoragePostgres(ctx, db)

	//db := memdb.New()
	//s := storage.NewStorageMemDb(db)
	handlers := handler.New(s)
	srv := new(Server)
	err = srv.Run(HTTP_PORT, handlers.InitRoutes())
	if err != nil {
		elog.Fatalf("error occured while running http server: %s", err.Error())
	}
}
