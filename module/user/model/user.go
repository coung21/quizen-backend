package model

import (
	"quizen/common"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	common.SQLModel
	Username  string        `json:"username" gorm:"column:username" validate:"required,min=6,max=100"`
	Email     string        `json:"email" gorm:"column:email" validate:"required,email"`
	Password  string        `json:"password,omitempty" gorm:"column:password" validate:"required,min=6,max=100"`
	Avatar    *common.Image `json:"avatar" gorm:"column:avatar"`
	IsVerifed bool          `json:"is_verified" gorm:"column:is_verified"`
}

func (User) TableName() string { return "users" }

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) Sanitize() {
	u.Password = ""
}
