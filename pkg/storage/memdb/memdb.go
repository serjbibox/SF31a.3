package memdb

import (
	"github.com/serjbibox/GoNews/pkg/models"
)

type DB []models.Post

func New() DB {
	return DB{}
}
