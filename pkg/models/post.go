package models

// Post - публикация.
type Post struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Content     string `json:"content" db:"content"`
	AuthorID    int    `json:"author_id" db:"author_id"`
	AuthorName  string `json:"name" db:"name"`
	CreatedAt   int64  `json:"created_at" db:"created_at"`
	PublishedAt int64  `json:"published_at" db:"published_at"`
}
