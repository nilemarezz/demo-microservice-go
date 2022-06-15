package service

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/nilemarezz/my-microservice/movie-service/internal/model"
	"github.com/nilemarezz/my-microservice/movie-service/internal/repository"
	"github.com/nilemarezz/my-microservice/movie-service/proto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetAll(t *testing.T) {
	repo := repository.NewMovieRepositoryMock()
	casts := []model.Cast{
		{Id: 1, Name: "test cast", Age: 60},
	}
	movies := []*model.Movie{
		{Id: 1, Name: "test", Description: "desc", ScreenDate: "01-01-1999", Cast: casts},
	}
	repo.On("All").Return(movies, nil)

	movieService := NewMovieService(repo)

	data, _ := movieService.GetAll(context.Background(), &proto.Empty{})

	assert.Equal(t, len(movies), len(data.Movies))
	assert.Equal(t, movies[0].Id, int(data.Movies[0].Id))
	assert.Equal(t, len(movies[0].Cast), len(data.Movies[0].Cast))

}

func TestGetAllError(t *testing.T) {
	repo := repository.NewMovieRepositoryMock()
	errExpect := errors.New("Error from db")
	repo.On("All").Return([]*model.Movie{}, errExpect)

	movieService := NewMovieService(repo)

	_, err := movieService.GetAll(context.Background(), &proto.Empty{})

	assert.ErrorIs(t, err, errExpect)

}

func TestGetByID(t *testing.T) {
	repo := repository.NewMovieRepositoryMock()
	casts := []model.Cast{
		{Id: 1, Name: "test cast", Age: 60},
	}
	movie := model.Movie{
		Id: 1, Name: "test", Description: "desc", ScreenDate: "01-01-1999", Cast: casts,
	}
	repo.On("ByID").Return(&movie, nil)

	movieService := NewMovieService(repo)

	data, _ := movieService.GetByID(context.Background(), &proto.MovieRequest{Id: 1})

	assert.Equal(t, 1, int(data.Movie.Id))

}

func TestGetByIDErrorNoRow(t *testing.T) {
	type TestCase struct {
		name      string
		exception error
		expect    error
	}

	otherException := errors.New("Error from db")
	cases := []TestCase{
		{name: "sql error no rows", exception: sql.ErrNoRows, expect: status.Error(codes.NotFound, "Item Not Found")},
		{name: "other exception", exception: otherException, expect: otherException},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			repo := repository.NewMovieRepositoryMock()
			repo.On("ByID").Return(&model.Movie{}, c.exception)
			movieService := NewMovieService(repo)

			_, err := movieService.GetByID(context.Background(), &proto.MovieRequest{Id: 1})
			assert.ErrorIs(t, err, c.expect)
		})
	}
}
