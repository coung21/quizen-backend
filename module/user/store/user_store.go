package store

import (
	"context"
	"quizen/common"
	"quizen/module/user/model"

	"gorm.io/gorm"
)

type UpdateUserParam struct {
	Password   string
	Avatar     *common.Image
	IsVerified bool
}

func (s UserStore) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s UserStore) GetUser(ctx context.Context, conditions map[string]interface{}, moreInfos ...string) (*model.User, error) {
	s.db.Begin()
	for _, info := range moreInfos {
		s.db = s.db.Preload(info)
	}

	var user model.User
	if err := s.db.Where(conditions).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.NotFound
		}
		return nil, err
	}
	return &user, nil
}

func (s UserStore) UpdateUser(ctx context.Context, conditions map[string]interface{}, user *model.User) (*model.User, error) {
	param := UpdateUserParam{
		Password:   user.Password,
		Avatar:     user.Avatar,
		IsVerified: user.IsVerifed,
	}
	if err := s.db.Model(&user).Where(conditions).Updates(param).Error; err != nil {
		return nil, err
	}

	return user, nil
}
