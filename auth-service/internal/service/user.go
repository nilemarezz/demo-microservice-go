package service

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/nilemarezz/my-microservice/auth-service/internal/repository"
	pb "github.com/nilemarezz/my-microservice/auth-service/proto"
	"go.opentelemetry.io/otel"
	otelCodes "go.opentelemetry.io/otel/codes"
)

type userService struct {
	pb.AuthServiceServer
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) pb.AuthServiceServer {
	return userService{repo: repo}
}

func (s userService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// start trace
	tracer := otel.GetTracerProvider().Tracer("auth-service")
	ctx, span := tracer.Start(ctx, "service/Login")
	defer span.End()

	user, err := s.repo.GetByUsername(ctx, req.Username)
	if len(user) <= 0 {
		return &pb.LoginResponse{Success: false, Message: "username not found", Token: ""}, nil
	}
	if err != nil {
		fmt.Println("err : ", err)
		if err == sql.ErrNoRows {
			res := &pb.LoginResponse{
				Success: false,
				Message: "User not fount",
				Token:   "",
			}
			return res, nil
		}
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		return nil, err
	}

	isMatch := checkPasswordHash(req.Password, user[0].Password)
	if !isMatch {
		res := &pb.LoginResponse{
			Success: false,
			Message: "Password not match",
			Token:   "",
		}
		return res, nil
	}

	// implement jwt

	token, err := generateToken(int(user[0].Id))
	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		return nil, err
	}

	return &pb.LoginResponse{
		Success: true,
		Token:   token,
	}, nil
}

func (s userService) Signup(ctx context.Context, req *pb.SignupRequest) (*pb.SignupResponse, error) {
	// start trace
	tracer := otel.GetTracerProvider().Tracer("auth-service")
	ctx, span := tracer.Start(ctx, "service/Login")
	defer span.End()

	users, err := s.repo.GetByUsername(ctx, req.Username)

	if err == nil || err == sql.ErrNoRows {
		if len(users) > 0 {
			return &pb.SignupResponse{Success: false, Username: "", Message: "username is exist"}, nil
		}
		// no user, save!!
		password, err := hashPassword(req.Password)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(otelCodes.Error, err.Error())
			return nil, err
		}
		err = s.repo.Add(ctx, &repository.User{Username: req.Username, Password: password})
		if err != nil {
			span.RecordError(err)
			span.SetStatus(otelCodes.Error, err.Error())
			return nil, err
		}
		return &pb.SignupResponse{Success: true, Username: req.Username, Message: "signup success"}, nil
	} else {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		return nil, err
	}
}

func (s userService) Authenticate(ctx context.Context, req *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
	tracer := otel.GetTracerProvider().Tracer("auth-service")
	_, span := tracer.Start(ctx, "service/Login")
	defer span.End()

	success, id, _ := isValidate(req.Token)
	if !success {
		return &pb.AuthenticateResponse{
			Success: success,
			Id:      0,
		}, nil
	}

	// success
	idInt, err := strconv.Atoi(id)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(otelCodes.Error, err.Error())
		return nil, err
	}
	res := &pb.AuthenticateResponse{
		Success: success,
		Id:      int32(idInt),
	}
	return res, nil
}
