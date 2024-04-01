package model

import "time"

type User struct {
	Id             int64
	Name           string
	HashedPassword []byte
	CreatedAt      time.Time
}
