package mock

import (
	"context"
	"quizen/module/flashcard/model"
)

type FlashCardStoreMock struct {
	CreateStudySetFn func(ctx context.Context, studySet *model.StudySet) (*model.StudySet, error)
	DeleteStudySetFn func(ctx context.Context, studySetID string) error
	UpdateStudySetFn func(ctx context.Context, studySet *model.StudySet) (model.StudySet, error)
	CreateCardsFn    func(ctx context.Context, card *model.Flashcard) (*model.Flashcard, error)
	UpdateCartFn     func(ctx context.Context, card *model.Flashcard) (*model.Flashcard, error)
}

func (m FlashCardStoreMock) CreateStudySet(ctx context.Context, studySet *model.StudySet) (*model.StudySet, error) {
	return m.CreateStudySetFn(ctx, studySet)
}

func (m FlashCardStoreMock) DeleteStudySet(ctx context.Context, studySetID string) error {
	return m.DeleteStudySetFn(ctx, studySetID)
}

func (m FlashCardStoreMock) UpdateStudySet(ctx context.Context, studySet *model.StudySet) (model.StudySet, error) {
	return m.UpdateStudySetFn(ctx, studySet)
}

func (m FlashCardStoreMock) CreateCards(ctx context.Context, card *model.Flashcard) (*model.Flashcard, error) {
	return m.CreateCardsFn(ctx, card)
}

func (m FlashCardStoreMock) UpdateCart(ctx context.Context, card *model.Flashcard) (*model.Flashcard, error) {
	return m.UpdateCartFn(ctx, card)
}
