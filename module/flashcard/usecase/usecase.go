package usecase

import (
	"context"
	"quizen/module/flashcard/model"
	"quizen/module/flashcard/store"
)

type UseCase interface {
	CreateStudySet(ctx context.Context, studySet *model.StudySet) (*model.StudySet, error)
}

type flashcardUseCase struct {
	store store.Store
}

func NewFlashcardUseCase(store store.Store) flashcardUseCase {
	return flashcardUseCase{store: store}
}
