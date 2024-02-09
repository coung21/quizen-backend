package mock

import (
	"context"
	"quizen/module/flashcard/model"
)

type FlashCardStoreMock struct {
	CreateStudySetFn func(ctx context.Context, studySet *model.StudySet) (*model.StudySet, error)
	DeleteStudySetFn func(ctx context.Context, studySetID string) error
}

func (m FlashCardStoreMock) CreateStudySet(ctx context.Context, studySet *model.StudySet) (*model.StudySet, error) {
	return m.CreateStudySetFn(ctx, studySet)
}

func (m FlashCardStoreMock) DeleteStudySet(ctx context.Context, studySetID string) error {
	return m.DeleteStudySetFn(ctx, studySetID)
}
