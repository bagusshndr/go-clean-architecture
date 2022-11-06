package mysql_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"

	"github.com/bxcodec/go-clean-arch/activity/repository"
	activityMysqlRepo "github.com/bxcodec/go-clean-arch/activity/repository/mysql"
	"github.com/bxcodec/go-clean-arch/domain"
)

func TestFetch(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mockArticles := []domain.Activity{
		{
			ID: 1, Email: "Bagus@gmail.com", Title: "title 1", UpdatedAt: time.Now(), CreatedAt: time.Now(),
		},
	}

	rows := sqlmock.NewRows([]string{"id", "email", "title", "updated_at", "created_at"}).
		AddRow(mockArticles[0].ID, mockArticles[0].Email, mockArticles[0].Title, mockArticles[0].UpdatedAt, mockArticles[0].CreatedAt)

	query := "SELECT id, email, title, updated_at, created_at FROM activity WHERE created_at > \\? ORDER BY created_at LIMIT \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := activityMysqlRepo.NewMysqlArticleRepository(db)
	cursor := repository.EncodeCursor(mockArticles[0].CreatedAt)
	num := int64(1)
	list, nextCursor, err := a.Fetch(context.TODO(), cursor, num)
	assert.NotEmpty(t, nextCursor)
	assert.NoError(t, err)
	assert.Len(t, list, 1)
}

func TestGetByID(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "email", "title", "updated_at", "created_at"}).
		AddRow(1, "title 1", "Content 1", time.Now(), time.Now())

	query := "SELECT id, email, title, updated_at, created_at FROM activity WHERE ID = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := activityMysqlRepo.NewMysqlArticleRepository(db)

	num := int64(5)
	anArticle, err := a.GetByID(context.TODO(), num)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestStore(t *testing.T) {
	now := time.Now()
	ar := &domain.Activity{
		ID:        12,
		Email:     "Content",
		Title:     "Judul",
		CreatedAt: now,
		UpdatedAt: now,
	}
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "INSERT activity SET email =?, title=?, updated_at=?, created_at=?"
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Email, ar.Title, ar.CreatedAt, ar.UpdatedAt).WillReturnResult(sqlmock.NewResult(12, 1))

	a := activityMysqlRepo.NewMysqlArticleRepository(db)

	err = a.Store(context.TODO(), ar)
	assert.NoError(t, err)
	assert.Equal(t, int64(12), ar.ID)
}

func TestGetByTitle(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"id", "email", "title", "updated_at", "created_at"}).
		AddRow(1, "title 1", "Content 1", time.Now(), time.Now())

	query := "SELECT id, email, title, updated_at, created_at FROM activity WHERE title = \\?"

	mock.ExpectQuery(query).WillReturnRows(rows)
	a := activityMysqlRepo.NewMysqlArticleRepository(db)

	title := "title 1"
	anArticle, err := a.GetByTitle(context.TODO(), title)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestDelete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "DELETE FROM activity WHERE id = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(12).WillReturnResult(sqlmock.NewResult(12, 1))

	a := activityMysqlRepo.NewMysqlArticleRepository(db)

	num := int64(12)
	err = a.Delete(context.TODO(), num)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	now := time.Now()
	ar := &domain.Activity{
		ID:        12,
		Email:     "Content",
		Title:     "Judul",
		CreatedAt: now,
		UpdatedAt: now,
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	query := "UPDATE activity set email=\\, title=\\?, updated_at=\\? WHERE ID = \\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(ar.Email, ar.Title, ar.UpdatedAt, ar.ID).WillReturnResult(sqlmock.NewResult(12, 1))

	a := activityMysqlRepo.NewMysqlArticleRepository(db)

	err = a.Update(context.TODO(), ar)
	assert.NoError(t, err)
}
