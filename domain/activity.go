package domain

import (
	"context"
	"time"
)

// Activity ...
type Activity struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email" validate:"required"`
	Title     string    `json:"title" validate:"required"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

// ArticleUsecase represent the article's usecases
type ActivityUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) ([]Activity, string, error)
	GetByID(ctx context.Context, id int64) (Activity, error)
	Update(ctx context.Context, ar *Activity) error
	GetByTitle(ctx context.Context, title string) (Activity, error)
	Store(context.Context, *Activity) error
	Delete(ctx context.Context, id int64) error
}

// ArticleRepository represent the article's repository contract
type ActivityRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Activity, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (Activity, error)
	GetByTitle(ctx context.Context, title string) (Activity, error)
	Update(ctx context.Context, ar *Activity) error
	Store(ctx context.Context, a *Activity) error
	Delete(ctx context.Context, id int64) error
}
