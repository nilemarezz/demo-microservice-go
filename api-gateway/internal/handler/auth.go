package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/nilemarezz/my-microservice/api-gateway/internal/model"
	"github.com/nilemarezz/my-microservice/api-gateway/internal/util"
	pb "github.com/nilemarezz/my-microservice/api-gateway/proto"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthHandler struct {
	client pb.AuthServiceClient
}

func NewAuthHandler(client pb.AuthServiceClient) AuthHandler {
	return AuthHandler{client}
}

// Login godoc
// @Summary      Login
// @Description  Login user
// @Tags         auth
// @Accept       json
// @Produce      json
// @param User body model.User true "User date to be login"
// @Success      202  {object}   model.LoginResponse
// @Failure      401  {object}  model.LoginResponse
// @Failure      404  {object}  util.JsonResponse
// @Failure      500  {object}  util.JsonResponse
// @Router       /auth/login [post]
func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// init span - tracing
	tracer := otel.GetTracerProvider().Tracer("http-Login")
	ctx, span := tracer.Start(r.Context(), "http-Login")
	defer span.End()
	// map json to model
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		util.ErrorJson(w, errors.New("request body invalid format"), http.StatusBadRequest)
	}
	// send gRPC
	span.SetAttributes(attribute.String("username", user.Username))
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := h.client.Login(ctx, &pb.LoginRequest{Username: user.Username, Password: user.Password})
	// check error from response
	if err != nil {
		if grpcError, ok := status.FromError(err); ok {
			util.ErrorJson(w, grpcError.Err(), http.StatusBadRequest)
			return
		} else {
			util.ErrorJson(w, err, http.StatusInternalServerError)
			return
		}
	}

	// check success from response
	if res.Success {
		util.WriteJson(w, http.StatusAccepted, model.LoginResponse{Token: res.Token, Message: res.Message})
	} else {
		util.WriteJson(w, http.StatusUnauthorized, model.LoginResponse{Token: "", Message: res.Message})
	}
}

// Signup godoc
// @Summary      Signup
// @Description  Signup user
// @Tags         auth
// @Accept       json
// @Produce      json
// @param User body model.User true "User date to be signup"
// @Success      202  {object}   model.SignupResponse
// @Failure      401  {object}  model.SignupResponse
// @Failure      404  {object}  util.JsonResponse
// @Failure      500  {object}  util.JsonResponse
// @Router       /auth/signup [post]
func (h AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	// init span - tracing
	tracer := otel.GetTracerProvider().Tracer("http-Signup")
	ctx, span := tracer.Start(r.Context(), "http-Signup")
	defer span.End()
	// map json to model
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		util.ErrorJson(w, errors.New("request body invalid format"), http.StatusBadRequest)
	}

	// send gRPC
	span.SetAttributes(attribute.String("username", user.Username))
	md := metadata.Pairs(
		"timestamp", time.Now().Format(time.StampNano),
	)
	ctx = metadata.NewOutgoingContext(ctx, md)
	res, err := h.client.Signup(ctx, &pb.SignupRequest{Username: user.Username, Password: user.Password})
	// check error from response
	if err != nil {
		if grpcError, ok := status.FromError(err); ok {
			util.ErrorJson(w, grpcError.Err(), http.StatusBadRequest)
			return
		} else {
			util.ErrorJson(w, err, http.StatusInternalServerError)
			return
		}
	}
	// check success from response
	if res.Success {
		util.WriteJson(w, http.StatusAccepted, model.SignupResponse{Success: res.Success, Username: res.Username, Message: res.Message})
	} else {
		util.WriteJson(w, http.StatusUnauthorized, model.SignupResponse{Success: res.Success, Username: "", Message: res.Message})
	}
}
