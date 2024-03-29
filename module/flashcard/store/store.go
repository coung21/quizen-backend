package store

import (
	"context"
	"quizen/module/flashcard/model"

	"gorm.io/gorm"
)

type Store interface {
	CreateStudySet(ctx context.Context, studySet *model.StudySet) (*model.StudySet, error)
	DeleteStudySet(ctx context.Context, studySetID string) error
	UpdateStudySet(ctx context.Context, studySet *model.StudySet) (model.StudySet, error)
	CreateCards(ctx context.Context, card *model.Flashcard) (*model.Flashcard, error)
	UpdateCart(ctx context.Context, card *model.Flashcard) (*model.Flashcard, error)
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
