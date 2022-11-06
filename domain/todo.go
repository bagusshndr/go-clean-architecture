package domain

import (
	"context"
	"time"
)

// Todo ...
type Todo struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title" validate:"required"`
	Content   string    `json:"content" validate:"required"`
	Author    Author    `json:"author"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// ArticleUsecase represent the article's usecases
type TodoUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Todo, string, error)
	GetByID(ctx context.Context, id int64) (Todo, error)
	Update(ctx context.Context, ar *Todo) error
	GetByTitle(ctx context.Context, title string) (Todo, error)
	Store(context.Context, *Todo) error
	Delete(ctx context.Context, id int64) error
}

// ArticleRepository represent the article's repository contract
type TodoRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Todo, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (Todo, error)
	GetByTitle(ctx context.Context, title string) (Todo, error)
	Update(ctx context.Context, ar *Todo) error
	Store(ctx context.Context, a *Todo) error
	Delete(ctx context.Context, id int64) error
}
