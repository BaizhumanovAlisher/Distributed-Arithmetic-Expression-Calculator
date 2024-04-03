package app

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type PassHasher struct {
	cost int
}

func NewPassHasher(cost int) *PassHasher {
	return &PassHasher{cost: cost}
}

func (p *PassHasher) GenerateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), p.cost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (p *PassHasher) Compare(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
