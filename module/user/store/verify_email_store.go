package store

import (
	"context"
	"quizen/module/user/model"
)

func (s *UserStore) CreateVerifyEmail(ctx context.Context, verifyEmail *model.VerifyEmail) (*model.VerifyEmail, error) {
	if err := s.db.Create(verifyEmail).Error; err != nil {
		return nil, err
	}
	return verifyEmail, nil
}

func (s *UserStore) UpdateVerifyEmail(ctx context.Context, email, secretCode string) (*model.VerifyEmail, error) {
	var verifyEmail model.VerifyEmail
	if err := s.db.Where("email = ? AND secret_code = ? AND is_used = ?", email, secretCode, false).First(&verifyEmail).Error; err != nil {
		return nil, err
	}
	return &verifyEmail, nil
}
