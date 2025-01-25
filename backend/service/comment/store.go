package comment

import (
	"database/sql"
	"github.com/Cleaach/threads/backend/types"
	"fmt"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetCommentsByThreadID(id int) ([]types.Comment, error) {
	rows, err := s.db.Query("SELECT * FROM comments WHERE thread_id = ?", id)
	if err != nil {
		return nil, err
	}

	comments := make([]types.Comment, 0)
	for rows.Next() {
		c, err := scanRowsIntoComment(rows)
		if err != nil {
			return nil, err
		}

		comments = append(comments, *c)
	}

	return comments, nil
}

func (s *Store) GetCommentByID(id int) (*types.Comment, error) {
	rows, err := s.db.Query("SELECT * FROM comments WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	c := new(types.Comment)
	for rows.Next() {
		c, err = scanRowsIntoComment(rows)
		if err != nil {
			return nil, err
		}
	}

	if c.ID == 0 {
		return nil, fmt.Errorf("user not found by id")
	}

	return c, nil
}

func scanRowsIntoComment(rows *sql.Rows) (*types.Comment, error) {
	comment := new(types.Comment)

	err := rows.Scan(
		&comment.ID,
		&comment.AuthorID,
		&comment.ThreadID,
		&comment.Content,
		&comment.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *Store) AddComment(comment types.Comment) error {
	_, err := s.db.Exec("INSERT INTO comments (author_id, thread_id, content) VALUES (?, ?, ?)", comment.AuthorID, comment.ThreadID, comment.Content)
	if err != nil {
		return err // Return the error if the query fails
	}

	return nil // Return nil if the insertion succeeds
}

func (s *Store) DeleteCommentByID(id int) error {
	_, err := s.db.Exec("DELETE FROM comments WHERE id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) EditComment(id int, comment types.Comment) error {
	_, err := s.db.Exec("UPDATE comments SET content = ? WHERE id = ?;", comment.Content, id)
	if err != nil {
		return err 
	}

	return nil 
}