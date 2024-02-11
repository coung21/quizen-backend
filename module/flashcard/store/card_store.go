package store

import (
	"context"
	"quizen/module/flashcard/model"
)

func (s flashcardStore) CreateCards(ctx context.Context, card *model.Flashcard) (*model.Flashcard, error) {
	if err := s.db.Create(card).Error; err != nil {
		return nil, err
	}
	// fmt.Println(card)
	return card, nil
}

func (s flashcardStore) UpdateCart(ctx context.Context, card *model.Flashcard) (*model.Flashcard, error) {
	var newCard model.Flashcard
	if err := s.db.Model(&model.Flashcard{}).Where("id = ?", card.ID).Updates(card).Scan(&newCard).Error; err != nil {
		return nil, err
	}

	// fmt.Println(newCard)

	return &newCard, nil
}
