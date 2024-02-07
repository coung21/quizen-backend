package store

import (
	"context"
	"quizen/module/user/model"
	"time"
)

func (s UserStore) CreateVerifyEmail(ctx context.Context, verifyEmail *model.VerifyEmail) (*model.VerifyEmail, error) {
	if err := s.db.Create(verifyEmail).Error; err != nil {
		return nil, err
	}
	return verifyEmail, nil
}

func (s UserStore) UpdateVerifyEmail(ctx context.Context, email, secretCode string) (*model.VerifyEmail, error) {
	var verifyEmail model.VerifyEmail
	verifyEmail.IsUsed = true
	verifyEmail.Email = email
	if err := s.db.Model(&verifyEmail).Where("email = ? AND secret_code = ? AND expired_at > ? AND is_used <> TRUE", email, secretCode, time.Now()).Update("is_used", true).Error; err != nil {
		return nil, err
	}
	return &verifyEmail, nil
}
