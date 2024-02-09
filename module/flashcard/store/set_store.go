package store

import (
	"context"
	"quizen/common"
	"quizen/module/flashcard/model"

	"gorm.io/gorm"
)

func (s flashcardStore) CreateStudySet(ctx context.Context, studySet *model.StudySet) (*model.StudySet, error) {
	if err := s.db.Create(studySet).Error; err != nil {
		return nil, err
	}
	return studySet, nil
}

func (s flashcardStore) DeleteStudySet(ctx context.Context, studySetID string) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := tx.Where("id = ?", studySetID).Delete(&model.StudySet{}).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return common.NotFound
		}
		return err
	}

	if err := tx.Where("study_set_id = ?", studySetID).Delete(&model.Flashcard{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
