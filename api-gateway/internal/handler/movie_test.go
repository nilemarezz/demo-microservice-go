package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/nilemarezz/my-microservice/api-gateway/internal/model"
	"github.com/nilemarezz/my-microservice/api-gateway/internal/util"
	pb "github.com/nilemarezz/my-microservice/api-gateway/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type movieServiceClientMock struct {
	mock.Mock
}

func NewMovieServiceClientMock() *movieServiceClientMock {
	return &movieServiceClientMock{}
}

func (m *movieServiceClientMock) GetAll(ctx context.Context, in *pb.Empty, opts ...grpc.CallOption) (*pb.MoviesResponse, error) {
	args := m.Called()
	return args.Get(0).(*pb.MoviesResponse), args.Error(1)
}

func (m *movieServiceClientMock) GetByID(ctx context.Context, in *pb.MovieRequest, opts ...grpc.CallOption) (*pb.MovieResponse, error) {
	args := m.Called()
	return args.Get(0).(*pb.MovieResponse), args.Error(1)
}

func (m *movieServiceClientMock) AddMovie(ctx context.Context, in *pb.Movie, opts ...grpc.CallOption) (*pb.AddMovieResponse, error) {
	args := m.Called()
	return args.Get(0).(*pb.AddMovieResponse), args.Error(1)
}

func TestGetMovies(t *testing.T) {
	client := NewMovieServiceClientMock()
	cast := &pb.Cast{
		Id:   1,
		Name: "test name",
		Age:  60,
	}
	movie := &pb.Movie{
		Id:          1,
		Name:        "test",
		Description: "desc",
		ScreenDate:  "01-01-1999",
		Cast:        []*pb.Cast{cast},
	}

	value := &pb.MoviesResponse{
		Movies: []*pb.Movie{movie},
	}

	client.On("GetAll").Return(value, nil)

	movieHandler := NewMovieHandler(client)

	w := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/movies/", movieHandler.GetMovies).Methods(http.MethodGet)
	r.ServeHTTP(w, httptest.NewRequest("GET", "/movies/", nil))

	res := []model.Movie{}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, http.StatusAccepted, w.Code)
	assert.Equal(t, 1, len(res))
}

func TestGetMoviesErrorgRPC(t *testing.T) {
	client := NewMovieServiceClientMock()

	client.On("GetAll").Return(&pb.MoviesResponse{}, status.Error(codes.NotFound, "Item Not Found"))

	movieHandler := NewMovieHandler(client)

	w := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/movies/", movieHandler.GetMovies).Methods(http.MethodGet)
	r.ServeHTTP(w, httptest.NewRequest("GET", "/movies/", nil))

	// unmarshall and check the data
	// unmarshall and check the data
	res := util.JsonResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.True(t, res.Error)

}

func TestGetMoviesError(t *testing.T) {
	client := NewMovieServiceClientMock()

	client.On("GetAll").Return(&pb.MoviesResponse{}, errors.New("error from calling service"))

	movieHandler := NewMovieHandler(client)

	w := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/movies/", movieHandler.GetMovies).Methods(http.MethodGet)
	r.ServeHTTP(w, httptest.NewRequest("GET", "/movies/", nil))

	// unmarshall and check the data
	res := util.JsonResponse{}
	err := json.Unmarshal(w.Body.Bytes(), &res)
	if err != nil {
		fmt.Println(err)
	}

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.True(t, res.Error)
	assert.Equal(t, "error from calling service", res.Message)

}
