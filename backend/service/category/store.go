package category

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

func (s *Store) GetCategoryNameByID(id int) (*types.Category, error) {
	rows, err := s.db.Query("SELECT * FROM categories WHERE id = ?", id)
	if err != nil {
		return nil, err
	}

	c := new(types.Category)
	for rows.Next() {
		c, err = scanRowIntoCategory(rows)
		if err != nil {
			return nil, err
		}
	}

	if c.ID == 0 {
		return nil, fmt.Errorf("category not found by id")
	}

	return c, nil
}

func (s *Store) GetCategoryIDByName(name string) (*types.Category, error) {
    rows, err := s.db.Query("SELECT * FROM categories WHERE name = ?", name)
    if err != nil {
        return nil, err
    }

    c := new(types.Category)
    for rows.Next() {
        c, err = scanRowIntoCategory(rows)
        if err != nil {
            return nil, err
        }
    }

    if c.ID == 0 {
        return nil, fmt.Errorf("category %s not found by name", name)
    }

    return c, nil
}


func scanRowIntoCategory(rows *sql.Rows) (*types.Category, error) {
    category := new(types.Category)

    err := rows.Scan(
        &category.ID,
        &category.Name,
    )

    if err != nil {
        return nil, err
    }

    return category, nil
}


func (s *Store) CreateCategory(category types.Category) error {
	_, err := s.db.Exec("INSERT INTO categories (name) VALUES (?)", category.Name)
	if err != nil {
		return err // Return the error if the query fails
	}

	return nil // Return nil if the insertion succeeds
}

func (s *Store) GetThreadsByCategoryID(categoryID int) ([]types.Thread, error) {
	rows, err := s.db.Query("SELECT * FROM threads WHERE category_id = ?", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var threads []types.Thread
	for rows.Next() {
		thread := new(types.Thread)
		err := rows.Scan(&thread.ID, &thread.AuthorID, &thread.CategoryID, &thread.Title, &thread.Content, &thread.CreatedAt)
		if err != nil {
			return nil, err
		}
		threads = append(threads, *thread)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return threads, nil
}
