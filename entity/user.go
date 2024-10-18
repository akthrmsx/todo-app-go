package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserID int64

type User struct {
	ID       UserID    `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
