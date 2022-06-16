package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/nilemarezz/my-microservice/api-gateway/internal/util"
	pb "github.com/nilemarezz/my-microservice/api-gateway/proto"
)

type authMiddleware struct {
	client pb.AuthServiceClient
}

func NewAuthMiddleware(client pb.AuthServiceClient) authMiddleware {
	return authMiddleware{client}
}

func (m authMiddleware) VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		authHeader := r.Header.Get("Authorization")
		if authHeader != "" {
			token := getTokenFromHeader(authHeader)
			fmt.Println(token)
			res, err := m.client.Authenticate(r.Context(), &pb.AuthenticateRequest{Token: token})
			fmt.Println("middleware -> auth_service : ", res)
			if err != nil {
				fmt.Println("error from middleware : ", err)
				util.ErrorJson(w, err, http.StatusUnauthorized)
				return
			}
			if res.Success {
				// context.Set(r, "user_id", res.Id)
				next.ServeHTTP(w, r)
			} else {
				util.ErrorJson(w, errors.New("Unauthorized"), http.StatusUnauthorized)
				return
			}
		} else {
			util.ErrorJson(w, errors.New("UnAuthorized"), http.StatusUnauthorized)
			return
		}
	})
}

func getTokenFromHeader(header string) string {
	splitToken := strings.Split(header, "Bearer")
	if len(splitToken) == 2 {
		return strings.TrimSpace(splitToken[1])
	}
	return ""
}
