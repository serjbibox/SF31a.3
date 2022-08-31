package main

import (
	"context"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/serjbibox/GoNews/pkg/handler"
	"github.com/serjbibox/GoNews/pkg/storage"
	"github.com/serjbibox/GoNews/pkg/storage/postgresql"
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
	connString, err := postgresql.GetConnectionString()
	if err != nil {
		elog.Fatal(err)
	}
	db, err := postgresql.New(connString)
	if err != nil {
		elog.Fatal(err)
	}
	defer db.Close()
	s, err := storage.NewStoragePostgres(ctx, db)
	if err != nil {
		elog.Fatal(err)
	}

	//db, err := mongodb.New(ctx)
	//if err != nil {
	//	elog.Fatal(err)
	//}
	//defer db.Disconnect(ctx)
	//s, err := storage.NewStorageMongodb(ctx, db)
	//if err != nil {
	//	elog.Fatal(err)
	//}

	//db, err := memdb.New()
	//if err != nil {
	//	elog.Fatal(err)
	//}
	//s, err := storage.NewStorageMemDb(db)
	//if err != nil {
	//	elog.Fatal(err)
	//}
	handlers, err := handler.New(s)
	if err != nil {
		elog.Fatal(err)
	}
	srv := new(Server)
	elog.Fatal(srv.Run(HTTP_PORT, handlers.InitRoutes()))
}
