package usecase

import "context"

func (uc flashcardUseCase) DeleteStudySet(ctx context.Context, studySetID string) error {
	return uc.store.DeleteStudySet(ctx, studySetID)
}
