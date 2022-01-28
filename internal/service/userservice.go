// Package service contain all business logic.
package service

import (
	"errors"
	"fmt"
	"github.com/Mykola-Mateichuk/golearn/internal/hasher"
	"github.com/Mykola-Mateichuk/golearn/internal/model"
	"github.com/Mykola-Mateichuk/golearn/internal/token"
	"time"
)

// Repository contain all methods to dial with user.
type Repository interface {
	AddUser(user model.User) (model.User, error)
	GetUsers() ([]model.User, error)
	GetUserByName(name string) (model.User, error)
}

// UserService contain repository link.
type UserService struct {
	repo Repository
	tokenMaker token.Maker
}

// NewUserService create user service.
func NewUserService(repo Repository, tokenMaker token.Maker) UserService {
	return UserService{
		repo: repo,
		tokenMaker: tokenMaker,
	}
}

// AddUser validate and added new user.
func (uservice UserService) AddUser (user model.User) (model.User, error) {
	err := uservice.ValidateNewUser(user)
	if err != nil {
		return user, err
	}

	// Add password hash.
	user.Password, _ = hasher.HashPassword(user.Password)
	if err != nil {
		return user, err
	}

	return uservice.repo.AddUser(user)
}

// GetUserList create list of all existing users.
func (uservice UserService) GetUserList () ([]model.User, error) {
	return uservice.repo.GetUsers()
}

// ValidateNewUser check if user is
func (uservice UserService) ValidateNewUser(user model.User) error {
	// @todo check user by other params not only by name.
	id, err := uservice.GetUserIdByName(user)
	if err != nil {
		return err
	}
	if id != "" {
		return errors.New(fmt.Sprintf("User with name: %s already exist", user.UserName))
	}

	return nil
}

// GetUserIdByName returns user id or empty string if user is not exist.
func (uservice UserService) GetUserIdByName(user model.User) (string, error) {
	var id string

	users, err := uservice.GetUserList()
	if err != nil {
		return "", err
	}

	for i := range users {
		isPasswordCorect := hasher.CheckPasswordHash(user.Password, users[i].Password)
		if err != nil {
			return id, err
		}

		if users[i].UserName == user.UserName && isPasswordCorect {
			id = users[i].Id
		}
	}

	return id, err
}

// GetUserByName returns user by name.
func (uservice UserService) GetUserByName(name string) (model.User, error) {
	user, err := uservice.repo.GetUserByName(name)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

// GetLoginLink create login link for user.
func (uservice UserService) GetLoginLink(user model.User) (string, error) {
	id, err := uservice.GetUserIdByName(user)
	if err != nil {
		return "", err
	}

	link := ""
	if id != "" {

		accessTokenStr, err := uservice.tokenMaker.CreateToken(user.UserName, 15*time.Minute)
		if err != nil {
			return "", errors.New("Can't create new token")
		}

		url := "/chat/ws.rtm.start?token=" + accessTokenStr
		link = fmt.Sprintf("You link to login into chat: %s", url)
	} else {
		return link, errors.New(fmt.Sprintf("Wrong user name or password"))
	}

	return link, err
}