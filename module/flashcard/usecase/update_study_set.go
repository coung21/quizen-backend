package usecase

import (
	"context"
	"fmt"
	"quizen/module/flashcard/model"

	"github.com/google/uuid"
)

func (uc flashcardUseCase) UpdateStudySet(ctx context.Context, studySet *model.StudySet) (*model.StudySet, error) {
	var updatedStudySet model.StudySet
	var err error

	if studySet.SetName != "" || studySet.Description != "" {
		updatedStudySet, err = uc.store.UpdateStudySet(ctx, studySet)
		if err != nil {
			return nil, err
		}
		updatedStudySet.Flashcards = studySet.Flashcards
	}

	for i, card := range updatedStudySet.Flashcards {
		if card.ID != uuid.Nil {
			newCard, err := uc.store.UpdateCart(ctx, &card)
			if err != nil {
				return nil, err
			}
			updatedStudySet.Flashcards[i] = *newCard
		} else {
			card.StudySetID = studySet.ID
			newCard, err := uc.store.CreateCards(ctx, &card)
			if err != nil {
				return nil, err
			}

			updatedStudySet.Flashcards[i] = *newCard
		}
	}

	fmt.Println(updatedStudySet.Flashcards)
	return &updatedStudySet, err
}
