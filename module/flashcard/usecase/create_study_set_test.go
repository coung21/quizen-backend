package usecase

import (
	"context"
	"quizen/common"
	"quizen/module/flashcard/mock"
	"quizen/module/flashcard/model"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateStudySet(t *testing.T) {
	now := time.Now()

	mockStore := mock.FlashCardStoreMock{
		CreateStudySetFn: func(ctx context.Context, studySet *model.StudySet) (*model.StudySet, error) {
			newStudySet := model.StudySet{
				SQLModel: common.SQLModel{
					ID:        uuid.New(),
					CreatedAt: &now,
					UpdatedAt: &now,
				},
				UserID:      studySet.UserID,
				SetName:     studySet.SetName,
				Description: studySet.Description,
				Flashcards:  studySet.Flashcards,
			}

			return &newStudySet, nil
		},
	}

	uc := NewFlashcardUseCase(mockStore)

	t.Run("success", func(t *testing.T) {
		studySet := model.StudySet{
			UserID:      uuid.New(),
			SetName:     "Test Set",
			Description: "Test Description",
			Flashcards: []model.Flashcard{
				{
					Term:       "Test Term",
					Definition: "Test Definition",
				},
				{
					Term:       "Test Term",
					Definition: "Test Definition",
				},
				{
					Term:       "Test Term",
					Definition: "Test Definition",
				},
				{
					Term:       "Test Term",
					Definition: "Test Definition",
				},
			},
		}

		got, err := uc.CreateStudySet(context.Background(), &studySet)

		assert.NoErrorf(t, err, "expected no error, got %v", err)
		assert.Equal(t, studySet.UserID, got.UserID)
		assert.Equal(t, studySet.SetName, got.SetName)
		assert.Equal(t, studySet.Description, got.Description)
	})

	t.Run("erorr flashcard length", func(t *testing.T) {
		studySet := model.StudySet{
			UserID:      uuid.New(),
			SetName:     "Test Set",
			Description: "Test Description",
			Flashcards:  []model.Flashcard{},
		}

		_, err := uc.CreateStudySet(context.Background(), &studySet)

		assert.Error(t, err)
		assert.Equal(t, model.ErrFlashCardLen, err)
	})
}
