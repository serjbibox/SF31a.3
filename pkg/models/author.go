package models

type Author struct {
	ID   string `json:"id" db:"id" bson:"_id"`
	Name string `json:"name" db:"name" bson:"name"`
}
