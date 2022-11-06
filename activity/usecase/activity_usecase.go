package usecase

import (
	"context"
	"time"

	"github.com/bxcodec/go-clean-arch/domain"
)

type activityUsecase struct {
	articleRepo    domain.ActivityRepository
	contextTimeout time.Duration
}

func NewArticleUsecase(a domain.ActivityRepository, timeout time.Duration) domain.ActivityUsecase {
	return &activityUsecase{
		articleRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *activityUsecase) Fetch(c context.Context, cursor string, num int64) (res []domain.Activity, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, nextCursor, err = a.articleRepo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	return
}

func (a *activityUsecase) GetByID(c context.Context, id int64) (res domain.Activity, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.articleRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (a *activityUsecase) Update(c context.Context, ar *domain.Activity) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return a.articleRepo.Update(ctx, ar)
}

func (a *activityUsecase) GetByTitle(c context.Context, title string) (res domain.Activity, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err = a.articleRepo.GetByTitle(ctx, title)
	if err != nil {
		return
	}

	return
}

func (a *activityUsecase) Store(c context.Context, m *domain.Activity) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, _ := a.GetByTitle(ctx, m.Title)
	if existedArticle != (domain.Activity{}) {
		return domain.ErrConflict
	}

	err = a.articleRepo.Store(ctx, m)
	return
}

func (a *activityUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, err := a.articleRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedArticle == (domain.Activity{}) {
		return domain.ErrNotFound
	}
	return a.articleRepo.Delete(ctx, id)
}
