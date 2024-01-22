package repository

import "gorm.io/gorm"

type Repository interface {
}

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepo {
	return &userRepo{DB: db}
}
