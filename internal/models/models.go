// Package models for server.
package models

// User define model for contain users.
type User struct {
	Id string
	UserName string
	Password string
}

// CreateUserRequest defines model for CreateUserRequest.
type CreateUserRequest struct {
	Password string `json:"Password"`
	UserName string `json:"userName"`
}

// CreateUserResponse defines model for CreateUserResponse.
type CreateUserResponse struct {
	Id       *string `json:"id,omitempty"`
	UserName *string `json:"userName,omitempty"`
}

// LoginUserRequest defines model for LoginUserRequest.
type LoginUserRequest struct {
	// The Password for login in clear text
	Password string `json:"Password"`

	// The user name for login
	UserName string `json:"userName"`
}

// LoginUserResonse defines model for LoginUserResonse.
type LoginUserResonse struct {
	// A url for websoket API with a one-time token for starting chat
	Url string `json:"url"`
}

// CreateUserJSONBody defines parameters for CreateUser.
type CreateUserJSONBody CreateUserRequest

// LoginUserJSONBody defines parameters for LoginUser.
type LoginUserJSONBody LoginUserRequest

// CreateUserJSONRequestBody defines body for CreateUser for application/json ContentType.
type CreateUserJSONRequestBody CreateUserJSONBody

// LoginUserJSONRequestBody defines body for LoginUser for application/json ContentType.
type LoginUserJSONRequestBody LoginUserJSONBody

