package models

import (
	"crypto/rand"
)

type Token struct {
	Id uint
	Secret []byte `gorm:"unique"`
}

func (t *Token) Generate() error {
	bytes := make([]byte, 32)
	_, err := rand.Read(bytes)
	if err != nil {
		return err
	}
	t.Secret = bytes

	return nil
}
