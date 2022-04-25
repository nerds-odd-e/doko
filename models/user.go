package models

import (
	"time"

	"github.com/go-pg/pg/v10/orm"
)

type User struct {
	tableName       struct{} `pg:"users,alias:users"`
	Id              int
	UserName        string `pg:"username"`
	CreatedAt       *time.Time
	UpdatedAt       *time.Time `json:",omitempty"`
	UpdatedByUserId int        `pg:"updated_by"`
}
type Users []User

func (users *Users) SelectAll(db orm.DB) error {
	err := db.Model(users).Order("id ASC").Select()
	return err
}

func (user *User) Insert(db orm.DB) error {
	_, err := db.Model(user).Insert()
	return err
}
