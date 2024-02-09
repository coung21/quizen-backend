package usecase

import (
	"context"
	"quizen/module/flashcard/mock"
	"testing"
)

func TestDeleteStudySet(t *testing.T) {
	mockStore := mock.FlashCardStoreMock{
		DeleteStudySetFn: func(ctx context.Context, studySetID string) error {
			return nil
		},
	}

	uc := NewFlashcardUseCase(mockStore)

	t.Run("success", func(t *testing.T) {

		err := uc.DeleteStudySet(context.Background(), "studySetID")
		if err != nil {
			t.Error("expected nil")
		}
	})
}
