package models

import "context"

type User struct {
	ID       int64
	Email    string
	PassHash []byte
}

type UserSaver interface {
	SaveUser(
		ctx context.Context,
		email string,
		passHash []byte,
	) (uid int64, err error)
}

type UserProvider interface {
	User(
		ctx context.Context,
		email string,
	) (User, error)
}
