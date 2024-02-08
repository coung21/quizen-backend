package mock

import (
	"context"
	"quizen/module/flashcard/model"
)

type FlashCardStoreMock struct {
	CreateStudySetFn func(ctx context.Context, studySet *model.StudySet) (*model.StudySet, error)
}

func (m FlashCardStoreMock) CreateStudySet(ctx context.Context, studySet *model.StudySet) (*model.StudySet, error) {
	return m.CreateStudySetFn(ctx, studySet)
}
