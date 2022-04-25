package services

import (
	"context"

	"github.com/go-pg/pg/v10"
	"tdd.com/v1/models"
)

type InsertUserDto struct {
	UserName string `json:"username"`
}

type UserResponse struct {
	Error string
	User  *models.User
}

func ListUsers(ctx context.Context, db *pg.DB) (models.Users, error) {
	var users models.Users
	err := users.SelectAll(db)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func InsertUser(ctx context.Context, db *pg.DB, insertUserDto InsertUserDto) error {
	user := models.User{
		UserName: insertUserDto.UserName,
	}

	err := user.Insert(db)
	if err != nil {
		return err
	}

	return nil
}
