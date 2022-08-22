package models

// Post - публикация.
type Post struct {
	ID          int    `json:"id" db:"id" bson:"-"`
	MongoID     string `bson:"_id"`
	Title       string `json:"title" db:"title" bson:"title"`
	Content     string `json:"content" db:"content" bson:"content"`
	AuthorID    int    `json:"author_id" db:"author_id" bson:"author_id"`
	AuthorName  string `json:"name"`
	CreatedAt   int64  `json:"created_at" db:"created_at" bson:"created_at"`
	PublishedAt int64  `json:"published_at" db:"published_at" bson:"published_at"`
}

type Author struct {
	ID      int
	MongoID string `json:"id" db:"id" bson:"_id"`
	Name    string `json:"name" db:"name" bson:"name"`
}
