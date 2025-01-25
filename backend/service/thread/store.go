package thread

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

func (s *Store) GetThreads() ([]types.Thread, error) {
	rows, err := s.db.Query("SELECT * FROM threads")
	if err != nil {
		return nil, err
	}

	threads := make([]types.Thread, 0)
	for rows.Next() {
		t, err := scanRowsIntoThread(rows)
		if err != nil {
			return nil, err
		}

		threads = append(threads, *t)
	}

	return threads, nil
}

func (s *Store) GetThreadByID(id int) (*types.Thread, error) {
	rows, err := s.db.Query("SELECT * FROM threads WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	t := new(types.Thread)
	for rows.Next() {
		t, err = scanRowsIntoThread(rows)
		if err != nil {
			return nil, err
		}
	}

	if t.ID == 0 {
		return nil, fmt.Errorf("user not found by id")
	}

	return t, nil
}

func scanRowsIntoThread(rows *sql.Rows) (*types.Thread, error) {
	thread := new(types.Thread)

	err := rows.Scan(
		&thread.ID,
		&thread.AuthorID,
		&thread.CategoryID,
		&thread.Title,
		&thread.Content,
		&thread.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return thread, nil
}

func (s *Store) CreateThread(thread types.Thread) error {
	_, err := s.db.Exec("INSERT INTO threads (author_id, category_id, title, content) VALUES (?, ?, ?, ?)", thread.AuthorID, thread.CategoryID, thread.Title, thread.Content)
	if err != nil {
		return err // Return the error if the query fails
	}

	return nil // Return nil if the insertion succeeds
}

func (s *Store) EditThread(id int, thread types.Thread) error {
	_, err := s.db.Exec("UPDATE threads SET category_id = ?, title = ?, content = ? WHERE id = ?;", thread.CategoryID, thread.Title, thread.Content, id)
	if err != nil {
		return err 
	}

	return nil 
}

func (s *Store) DeleteThreadByID(id int) error {
	
	_, err := s.db.Exec("DELETE FROM comments WHERE thread_id = ?;", id)
	if err != nil {
		return err
	}
	
	
	_, err = s.db.Exec("DELETE FROM threads WHERE id = ?;", id)
	if err != nil {
		return err
	}

	return nil
}
