package usecase

import (
	"context"
	"quizen/module/flashcard/model"
)

func (uc flashcardUseCase) CreateStudySet(ctx context.Context, studySet *model.StudySet) (*model.StudySet, error) {
	if len(studySet.Flashcards) < 4 {
		return nil, model.ErrFlashCardLen
	}

	return uc.store.CreateStudySet(ctx, studySet)
}
