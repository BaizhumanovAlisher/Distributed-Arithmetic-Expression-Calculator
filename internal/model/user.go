package model

import "time"

type User struct {
	Id             int
	Name           string
	HashedPassword []byte
	Created        time.Time
}
