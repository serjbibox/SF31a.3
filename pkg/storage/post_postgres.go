package storage

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/serjbibox/GoNews/pkg/models"
)

type PostPostgres struct {
	db  *pgxpool.Pool
	ctx context.Context
}

func newPostPostgres(ctx context.Context, db *pgxpool.Pool) Post {
	return &PostPostgres{
		db:  db,
		ctx: ctx,
	}
}

func (s *PostPostgres) GetAll() ([]models.Post, error) {
	var posts []models.Post
	query := `
		SELECT 
			posts.id, 
			posts.title, 
			posts.content, 
			posts.author_id, 
			posts.created_at,
			authors.name
		FROM posts, authors
		WHERE posts.author_id = authors.id
		ORDER BY id;`
	rows, err := s.db.Query(s.ctx, query)
	if err != nil {
		return nil, err
	}
	var id, author_id int
	for rows.Next() {
		var p models.Post
		err = rows.Scan(
			&id,
			&p.Title,
			&p.Content,
			&author_id,
			&p.CreatedAt,
			&p.AuthorName,
		)
		if err != nil {
			return nil, err
		}
		p.ID = strconv.Itoa(id)
		p.AuthorID = strconv.Itoa(author_id)
		posts = append(posts, p)

	}
	return posts, rows.Err()
}

func (s *PostPostgres) Create(p models.Post) (string, error) {
	tx, err := s.db.Begin(s.ctx)
	if err != nil {
		return "", err
	}

	createPostQuery := `
		INSERT INTO 
			posts (title, content, author_id, created_at)
		VALUES ($1, $2, $3, $4) RETURNING id;`
	author_id, err := strconv.Atoi(p.AuthorID)
	if err != nil {
		return "", err
	}
	p.CreatedAt = time.Now().Unix()
	row := tx.QueryRow(s.ctx, createPostQuery,
		p.Title,
		p.Content,
		author_id,
		p.CreatedAt,
	)
	var id int
	err = row.Scan(&id)
	if err != nil {
		tx.Rollback(s.ctx)
		return "", err
	}
	return strconv.Itoa(id), tx.Commit(s.ctx)
}

func (s *PostPostgres) Update(p models.Post) error {
	id, err := strconv.Atoi(p.ID)
	if err != nil {
		return err
	}
	author_id, err := strconv.Atoi(p.AuthorID)
	if err != nil {
		return err
	}
	err = s.db.QueryRow(s.ctx, `
	UPDATE posts
	SET title = $1, 
		content = $2, 
		author_id = $3 
	WHERE id = $4	
	RETURNING id;
	`,
		p.Title,
		p.Content,
		author_id,
		id,
	).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
func (s *PostPostgres) Delete(id string) error {
	delId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	res, err := s.db.Exec(s.ctx, `
	DELETE FROM posts 
	WHERE id = $1	
	`,
		delId,
	)
	if err != nil {
		return err
	}
	if res.Delete() {
		if res.String() == "DELETE 0" {
			return errors.New("no post to delete")
		}
	}
	return nil
}
