package memdb

import (
	"github.com/serjbibox/GoNews/pkg/models"
)

type DB []models.Task

func NewMemDb() DB {
	return DB{}
}
