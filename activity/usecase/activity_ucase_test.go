package usecase_test

import (
	"context"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	ucase "github.com/bxcodec/go-clean-arch/activity/usecase"
	"github.com/bxcodec/go-clean-arch/domain"
	"github.com/bxcodec/go-clean-arch/domain/mocks"
)

func TestFetch(t *testing.T) {
	mockArticleRepo := new(mocks.ActivityRepositoryMock)
	mockArticle := domain.Activity{
		Title: "Hello",
		Email: "Content",
	}

	mockListArtilce := make([]domain.Activity, 0)
	mockListArtilce = append(mockListArtilce, mockArticle)
	log.Println("asu", mockListArtilce)
	t.Run("success", func(t *testing.T) {
		mockArticleRepo.On("Fetch", mock.Anything, mock.AnythingOfType("string"),
			mock.AnythingOfType("int64")).Return(mockListArtilce, "next-cursor", nil).Once()
		u := ucase.NewArticleUsecase(mockArticleRepo, time.Second*2)
		num := int64(1)
		cursor := "12"
		list, nextCursor, err := u.Fetch(context.TODO(), cursor, num)
		cursorExpected := "next-cursor"
		assert.Equal(t, cursorExpected, nextCursor)
		assert.NotEmpty(t, nextCursor)
		assert.NoError(t, err)
		assert.Len(t, list, len(mockListArtilce))

		mockArticleRepo.AssertExpectations(t)
	})

}

func TestGetByID(t *testing.T) {
	mockArticleRepo := new(mocks.ActivityRepositoryMock)
	mockArticle := domain.Activity{
		Title: "Hello",
		Email: "Content",
	}

	t.Run("success", func(t *testing.T) {
		mockArticleRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockArticle, nil).Once()
		u := ucase.NewArticleUsecase(mockArticleRepo, time.Second*2)

		a, err := u.GetByID(context.TODO(), mockArticle.ID)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockArticleRepo.AssertExpectations(t)
	})

}

func TestStore(t *testing.T) {
	mockArticleRepo := new(mocks.ActivityRepositoryMock)
	mockArticle := domain.Activity{
		Title: "Hello",
		Email: "Content",
	}

	t.Run("success", func(t *testing.T) {
		tempMockArticle := mockArticle
		tempMockArticle.ID = 0
		mockArticleRepo.On("GetByTitle", mock.Anything, mock.AnythingOfType("string")).Return(domain.Activity{}, domain.ErrNotFound).Once()
		mockArticleRepo.On("Store", mock.Anything, mock.AnythingOfType("*domain.Activity")).Return(nil).Once()

		u := ucase.NewArticleUsecase(mockArticleRepo, time.Second*2)

		err := u.Store(context.TODO(), &tempMockArticle)

		assert.NoError(t, err)
		assert.Equal(t, mockArticle.Title, tempMockArticle.Title)
		mockArticleRepo.AssertExpectations(t)
	})

}

func TestDelete(t *testing.T) {
	mockArticleRepo := new(mocks.ActivityRepositoryMock)
	mockArticle := domain.Activity{
		Title: "Hello",
		Email: "Content",
	}

	t.Run("success", func(t *testing.T) {
		mockArticleRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(mockArticle, nil).Once()

		mockArticleRepo.On("Delete", mock.Anything, mock.AnythingOfType("int64")).Return(nil).Once()

		u := ucase.NewArticleUsecase(mockArticleRepo, time.Second*2)

		err := u.Delete(context.TODO(), mockArticle.ID)

		assert.NoError(t, err)
		mockArticleRepo.AssertExpectations(t)
	})
	t.Run("article-is-not-exist", func(t *testing.T) {
		mockArticleRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{}, nil).Once()

		u := ucase.NewArticleUsecase(mockArticleRepo, time.Second*2)

		err := u.Delete(context.TODO(), mockArticle.ID)

		assert.Error(t, err)
		mockArticleRepo.AssertExpectations(t)
	})
	t.Run("error-happens-in-db", func(t *testing.T) {
		mockArticleRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int64")).Return(domain.Activity{}, errors.New("Unexpected Error")).Once()

		u := ucase.NewArticleUsecase(mockArticleRepo, time.Second*2)

		err := u.Delete(context.TODO(), mockArticle.ID)

		assert.Error(t, err)
		mockArticleRepo.AssertExpectations(t)
	})

}

func TestUpdate(t *testing.T) {
	mockArticleRepo := new(mocks.ActivityRepositoryMock)
	mockArticle := domain.Activity{
		Title: "Hello",
		Email: "Content",
		ID:    23,
	}

	t.Run("success", func(t *testing.T) {
		mockArticleRepo.On("Update", mock.Anything, &mockArticle).Once().Return(nil)

		u := ucase.NewArticleUsecase(mockArticleRepo, time.Second*2)

		err := u.Update(context.TODO(), &mockArticle)
		assert.NoError(t, err)
		mockArticleRepo.AssertExpectations(t)
	})
}
