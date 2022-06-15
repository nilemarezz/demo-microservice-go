package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

type User struct {
	Id       int32  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

type UserRepository interface {
	GetByUsername(context.Context, string) ([]*User, error)
	Add(context.Context, *User) error
}

type userRepository struct {
	DB *sqlx.DB
}

func NewMovieRepository(db *sqlx.DB) UserRepository {
	return userRepository{db}
}

func (r userRepository) GetByUsername(ctx context.Context, username string) ([]*User, error) {
	tracer := otel.GetTracerProvider().Tracer("auth-service")
	_, span := tracer.Start(ctx, "database")
	span.SetAttributes(attribute.String("username", username))
	defer span.End()

	sql := "SELECT id, username, password FROM users WHERE username = ?"
	var user []*User
	err := r.DB.SelectContext(ctx, &user, sql, username)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (r userRepository) Add(ctx context.Context, user *User) error {
	tracer := otel.GetTracerProvider().Tracer("auth-service")
	_, span := tracer.Start(ctx, "database")
	defer span.End()

	_, err := r.DB.NamedExecContext(ctx, `INSERT INTO users (username,password) VALUES (:username,:password)`,
		map[string]interface{}{
			"username": user.Username,
			"password": user.Password,
		})

	if err != nil {
		return err
	}
	return nil
}
