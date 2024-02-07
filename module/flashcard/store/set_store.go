package store

import (
	"context"
	"quizen/module/flashcard/model"
)

func (s flashcardStore) CreateStudySet(ctx context.Context, studySet *model.StudySet) (*model.StudySet, error) {
	if err := s.db.Create(studySet).Error; err != nil {
		return nil, err
	}
	return studySet, nil
}
