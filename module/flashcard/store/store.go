package store

import (
	"context"
	"quizen/module/flashcard/model"

	"gorm.io/gorm"
)

type Store interface {
	CreateStudySet(ctx context.Context, studySet *model.StudySet) (*model.StudySet, error)
}

type flashcardStore struct {
	db *gorm.DB
}

func NewFlashcardStore(db *gorm.DB) flashcardStore {
	return flashcardStore{db: db}
}

func (s flashcardStore) WithTransaction(ctx context.Context, txFunc func(ctx context.Context, tx Store) error) error {
	tx := s.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	txStore := flashcardStore{db: tx}

	err := txFunc(ctx, txStore)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
