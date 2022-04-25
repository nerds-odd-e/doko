package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectAllUsers(t *testing.T) {
	err := seedUsers()
	assert.NoError(t, err)

	defer tearDownAll()

	var users Users
	err = users.SelectAll(db)
	assert.NoError(t, err)
	assert.Equal(t, 3, len(users))
}

func TestInsertUser(t *testing.T) {
	defer tearDownAll()

	user := User{
		UserName: "Test",
	}
	err := user.Insert(db)
	assert.NoError(t, err)
	// TODO: Verify the user has been inserted
}

func seedUsers() error {
	users := Users{
		{UserName: "Tim"},
		{UserName: "John"},
		{UserName: "Bell"},
	}

	for _, item := range users {
		_, err := db.Model(&item).Insert()
		if err != nil {
			return err
		}
	}
	return nil
}

func tearDownAll() error {
	_, err := db.Exec("TRUNCATE users RESTART IDENTITY CASCADE;")
	return err
}
