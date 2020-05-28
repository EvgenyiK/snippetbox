package models

import (
	"errors"
	"time"
)

//ErrNoRecord error
var ErrNoRecord = errors.New("models: no matching record found")
//ErrInvalidCredentials error
var ErrInvalidCredentials = errors.New("models: invalid credentials")
//ErrDuplicateEmail error
var ErrDuplicateEmail = errors.New("models: duplicate email")

//Snippet type
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// User type with the columns in the database `users` table
type User struct{
	ID		int
	Name	string
	Email	string
	HashedPassword	[]byte
	Created	time.Time
	Active	bool
}
