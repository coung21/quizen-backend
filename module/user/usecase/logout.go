package usecase

import "context"

func (u userUsecase) Logout(ctx context.Context, sessionID string) error {
	return u.userStore.DeleteSession(ctx, map[string]interface{}{"id": sessionID})
}
