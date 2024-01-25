package store

import (
	"context"
	"quizen/common"
	"quizen/module/user/model"

	"gorm.io/gorm"
)

type UpdateUserParam struct {
	Password   *string
	Avatar     *common.Image
	IsVerified *bool
}

func (s *UserStore) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserStore) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.NotFound
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) GetUserById(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.NotFound
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserStore) UpdateUser(ctx context.Context, id int, user *model.User) (*model.User, error) {
	param := UpdateUserParam{
		Password:   &user.Password,
		Avatar:     user.Avatar,
		IsVerified: &user.IsVerifed,
	}
	if err := s.db.Model(&user).Where("id = ?", id).Updates(param).Error; err != nil {
		return nil, err
	}
	return user, nil
}
