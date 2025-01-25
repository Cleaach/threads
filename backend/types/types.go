package types

import "time"

type UserStore interface {
	GetUserByUsername(username string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(User) error
}

type ThreadStore interface {
	GetThreads() ([]Thread, error)
	CreateThread(Thread) error
	GetThreadByID(id int) (*Thread, error)
	DeleteThreadByID(id int) error
	EditThread(id int, thread Thread) error
}

type CategoryStore interface {
	GetCategoryIDByName(name string) (*Category, error)
	CreateCategory(category Category) error
	GetThreadsByCategoryID(categoryID int) ([]Thread, error)
	GetCategoryNameByID(id int) (*Category, error)
}

type CommentStore interface {
	GetCommentByID(id int) (*Comment, error)
	GetCommentsByThreadID(id int) ([]Comment, error)
	AddComment(Comment) error
	DeleteCommentByID(id int) error
	EditComment(id int, comment Comment) error
}


type Thread struct {
	ID int `json:"id"`
	AuthorID int `json:"authorId"`
	CategoryID int `json:"categoryId`
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type Category struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID int `json:"id"`
	AuthorID int `json:"authorId"`
	ThreadID int `json:"threadId"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"createdAt`
}

type RegisterUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateThreadPayload struct {
	Category string `json:"category"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

type CreateCommentPayload struct {
	AuthorID int `json:"authorId"`
	ThreadID int `json:"threadId"`
	Content string `json:"content"`
}


