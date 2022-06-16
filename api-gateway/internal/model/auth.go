package model

// User model info
// @Description User information
// @Description with username,password
type User struct {
	Username string `json:"username" example:"Hello1"`
	Password string `json:"password" example:"Hello"`
}

// LoginResponse model info
// @Description LoginResponse information
// @Description with token,message
type LoginResponse struct {
	Token   string `json:"token" example:"asdg123rgfsd3fs51"`
	Message string `json:"message" example:"login success"`
}

// SignupResponse model info
// @Description SignupResponse information
// @Description with success,username,message
type SignupResponse struct {
	Success  bool   `json:"success" example:"true"`
	Username string `json:"username" example:"username"`
	Message  string `json:"message" example:"signup success"`
}
