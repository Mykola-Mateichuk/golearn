// Package repository store data.
package repository

import "github.com/Mykola-Mateichuk/golearn/internal/model"

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

// GetUsers return all users.
func (mstorage *MemoryUserStorage) GetUsers() ([]model.User, error) {
	return mstorage.Users, nil
}