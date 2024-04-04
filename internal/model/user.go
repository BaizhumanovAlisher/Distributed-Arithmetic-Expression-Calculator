package model

import "time"

type User struct {
	Id             int64
	Name           string
	HashedPassword string
	CreatedAt      time.Time
}

func NewUser(name string, hashedPassword string) *User {
	return &User{Name: name, HashedPassword: hashedPassword}
}

type UserCredentials struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required"`
}
