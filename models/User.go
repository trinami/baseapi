package models

import "time"

type User struct {
	Id uint `gorm:"primary_key;auto_increment"`
	Username string `gorm:"unique"`
	Password string
	TokenId uint
	Token Token
	LastSeen time.Time
	PublicKeys []PublicKey `gorm:"foreignkey:UserId"`
}
