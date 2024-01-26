package usecase

import (
	"context"
	"quizen/module/user/model"
	"quizen/module/user/store"
)

func (u userUsecase) VerifyEmail(ctx context.Context, email, secretCode string) (*model.VerifyEmail, error) {
	var verifyEmail *model.VerifyEmail
	err := u.userStore.WithTransaction(ctx, func(ctx context.Context, tx store.Store) error {
		var err error
		_, err = u.userStore.UpdateUser(ctx, map[string]interface{}{"email": email}, &model.User{IsVerifed: true})
		if err != nil {
			return err
		}

		verifyEmail, err = u.userStore.UpdateVerifyEmail(ctx, email, secretCode)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return verifyEmail, nil
}
