package repository

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/nilemarezz/my-microservice/movie-service/internal/model"
	"go.opentelemetry.io/otel"
)

type MovieRepository interface {
	All(context.Context) ([]*model.Movie, error)
	ByID(uint32, context.Context) (*model.Movie, error)
	Add(model.Movie) error
}

type movieRepository struct {
	DB *sqlx.DB
}

func NewMovieRepository(db *sqlx.DB) MovieRepository {
	return movieRepository{db}
}

func (r movieRepository) All(ctx context.Context) ([]*model.Movie, error) {

	tracer := otel.GetTracerProvider().Tracer("movie-service")
	ctx, span := tracer.Start(ctx, "database")
	defer span.End()

	sql := `SELECT m.movie_id , m.name , m.description  , m.screen_date FROM movies m `
	var movies []*model.Movie
	err := r.DB.SelectContext(ctx, &movies, sql)
	if err != nil {
		return nil, err
	}
	for i, movie := range movies {
		cast, err := r.getCastByMovieID(uint(movies[i].Id), ctx)
		fmt.Println(cast)
		if err != nil {
			return nil, err
		}
		movie.Cast = cast
	}
	return movies, nil
}

func (r movieRepository) ByID(id uint32, ctx context.Context) (*model.Movie, error) {

	tracer := otel.GetTracerProvider().Tracer("movie-service")
	ctx, span := tracer.Start(ctx, "database")
	defer span.End()

	sql := `SELECT m.movie_id , m.name , m.description  , m.screen_date FROM movies m WHERE m.movie_id = ?`
	var movie model.Movie
	err := r.DB.GetContext(ctx, &movie, sql, id)
	if err != nil {
		return nil, err
	}
	cast, err := r.getCastByMovieID(uint(movie.Id), ctx)
	if err != nil {
		return nil, err
	}
	movie.Cast = cast

	return &movie, nil
}

func (r movieRepository) Add(model.Movie) error {
	panic("")
}

func (r movieRepository) getCastByMovieID(movieID uint, ctx context.Context) ([]model.Cast, error) {
	sql := `
	SELECT c.name , c.age , c.id 
	FROM movies m 
	JOIN movie_celebritry mc ON m.movie_id = mc.movie_id 
	JOIN celebrities c ON c.id  = mc.celebritry_id
	WHERE m.movie_id = ?`
	var casts []model.Cast
	err := r.DB.SelectContext(ctx, &casts, sql, movieID)
	if err != nil {
		return nil, err
	}
	return casts, nil
}
