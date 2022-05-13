package services

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"tdd.com/v1/models"
)

func XTestPopulatingA2DArray(t *testing.T) {
	const size = int(100)
	a := [size][size]int{}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {		
			a[i][j] = i * j
		}
	}
	assert.Equal(t, (size-1)*(size-1), a[size-1][size-1])
}

func TestListUsers(t *testing.T) {
	ctx := context.Background()

	// Arrange
	seedUsers()
	defer tearDownAll()

	// Act
	users, err := ListUsers(ctx, db)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 3, len(users))
}

func TestInsertUser(t *testing.T) {
	ctx := context.Background()

	// Arrange
	defer tearDownAll()

	insertUserDto := InsertUserDto{
		UserName: "Alex",
	}

	// Act
	err := InsertUser(ctx, db, insertUserDto)
	assert.NoError(t, err)

	// Assert that Count of Users increase by 1
	users, err := ListUsers(ctx, db)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(users))
}

func seedUsers() {
	users := models.Users{
		{UserName: "Tim"},
		{UserName: "John"},
		{UserName: "Bell"},
	}

	for _, item := range users {
		db.Model(&item).Insert()
	}
}

func tearDownAll() {
	db.Exec("TRUNCATE users RESTART IDENTITY CASCADE;")
}
