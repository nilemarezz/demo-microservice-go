package repository

import (
	"context"

	"github.com/nilemarezz/my-microservice/movie-service/internal/model"
	"github.com/stretchr/testify/mock"
)

type movieRepositoryMock struct {
	mock.Mock
}

func NewMovieRepositoryMock() *movieRepositoryMock {
	return &movieRepositoryMock{}
}

func (m movieRepositoryMock) All(ctx context.Context) ([]*model.Movie, error) {
	args := m.Called()
	return args.Get(0).([]*model.Movie), args.Error(1)
}

func (m movieRepositoryMock) ByID(id uint32, ctx context.Context) (*model.Movie, error) {
	args := m.Called()
	return args.Get(0).(*model.Movie), args.Error(1)
}

func (m movieRepositoryMock) Add(model.Movie) error {
	panic("")
}
