package mocks

import (
	context "context"

	domain "github.com/bxcodec/go-clean-arch/domain"
	mock "github.com/stretchr/testify/mock"
)

type ActivityRepositoryMock struct {
	mock.Mock
}

func (m *ActivityRepositoryMock) Delete(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)

	return args.Error(0)
}

func (m *ActivityRepositoryMock) Fetch(ctx context.Context, cursor string, num int64) ([]domain.Activity, string, error) {
	args := m.Called(ctx, cursor, num)

	return args.Get(0).([]domain.Activity), args.Get(0).(string), args.Error(1)
}

func (m *ActivityRepositoryMock) GetByID(ctx context.Context, id int64) (domain.Activity, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(domain.Activity), args.Error(1)
}

func (m *ActivityRepositoryMock) GetByTitle(ctx context.Context, title string) (domain.Activity, error) {
	args := m.Called(ctx, title)

	return args.Get(0).(domain.Activity), args.Error(1)
}

func (m *ActivityRepositoryMock) Store(_a0 context.Context, _a1 *domain.Activity) error {
	args := m.Called(_a0, _a1)

	return args.Error(0)
}

func (m *ActivityRepositoryMock) Update(ctx context.Context, ar *domain.Activity) error {
	args := m.Called(ctx, ar)

	return args.Error(0)
}
