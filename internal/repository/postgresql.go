// Package repository store data.
package repository

import (
	"database/sql"
	"fmt"
	"github.com/Mykola-Mateichuk/golearn/internal/model"
)

// PostgreSqlUserStorage save users.
type PostgreSqlUserStorage struct {
	DB *sql.DB
}

// NewPostgreSqlStorage construct new object.
func NewPostgreSqlStorage(db *sql.DB) *PostgreSqlUserStorage {
	return &PostgreSqlUserStorage{DB: db}
}

// AddUser append new user to storage.
func (mstorage *PostgreSqlUserStorage) AddUser(user model.User) (model.User, error) {
	//mstorage.Users = append(mstorage.Users, user)
	//query, err := mstorage.DB.Query("INSERT INTO entries (guestName, content) values ('first guest', 'I got here!');")
	//if err != nil {
	//	return model.User{}, err
	//}
	_, err := mstorage.DB.Exec("INSERT INTO users (name, password) values ($1, $2)", user.UserName, user.Password)
	if err != nil{
		return user, err
	}
	return user, nil
}

// GetUsers return all users.
func (mstorage *PostgreSqlUserStorage) GetUsers() ([]model.User, error) {
	var users []model.User

	rows, err := mstorage.DB.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}

	for rows.Next() {
		user := model.User{}
		err := rows.Scan(&user.Id, &user.UserName, &user.Password)
		if err != nil{
			fmt.Println(err)
			continue
		}
		users = append(users, user)
	}

	return users, nil
}