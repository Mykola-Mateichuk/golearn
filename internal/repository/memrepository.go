// Package repository store data.
package repository

import (
	"fmt"
	"github.com/Mykola-Mateichuk/golearn/internal/model"
)

// MemoryUserStorage Contain all users.
type MemoryUserStorage struct {
	Users []model.User
}

// NewMemoryStorage construct new object.
func NewMemoryStorage() *MemoryUserStorage {
	var users []model.User
	return &MemoryUserStorage{Users: users}
}

// AddUser append new user to storage.
func (mstorage *MemoryUserStorage) AddUser(user model.User) (model.User, error) {
	mstorage.Users = append(mstorage.Users, user)
	return user, nil
}

// GetUserByName return all users.
func (mstorage *MemoryUserStorage) GetUserByName(name string) (model.User, error) {
	users, err := mstorage.GetUsers()

	if err != nil {
		fmt.Println(err)
		return model.User{}, err
	}

	for _, user := range users {
		if user.UserName == name {
			return user, nil
		}
	}

	return model.User{}, nil
}

// GetUsers return all users.
func (mstorage *MemoryUserStorage) GetUsers() ([]model.User, error) {
	return mstorage.Users, nil
}