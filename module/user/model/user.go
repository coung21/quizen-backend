package model

import (
	"quizen/common"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	common.SQLModel
	Username string        `json:"username" gorm:"column:username,unique,not null,type:varchar(100)" validate:"required,min=6,max=100"`
	Email    string        `json:"email" gorm:"column:email,unique,not null,type:varchar(100)" validate:"required,email"`
	Password string        `json:"password,omitempty" gorm:"column:password" validate:"required,min=6,max=100"`
	Avatar   *common.Image `json:"avatar" gorm:"column:avatar"`
	IsVerify bool          `json:"is_verify" gorm:"column:is_verify,default:false"`
}

func (User) TableName() string { return "user" }

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
