package store

import (
	"context"
	"quizen/common"
	"quizen/module/user/model"

	"gorm.io/gorm"
)

func (s UserStore) CreateSession(ctx context.Context, session *model.Session) (*model.Session, error) {
	if err := s.db.Create(session).Error; err != nil {
		return nil, err
	}
	return session, nil
}

func (s UserStore) GetSession(ctx context.Context, conditions map[string]interface{}, moreInfos ...string) (*model.Session, error) {
	for _, info := range moreInfos {
		s.db = s.db.Preload(info)
	}

	var session model.Session
	if err := s.db.Where(conditions).First(&session).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.NotFound
		}
		return nil, err
	}

	return &session, nil

}

func (s UserStore) DeleteSession(ctx context.Context, conditions map[string]interface{}) error {
	if err := s.db.Where(conditions).Delete(&model.Session{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.NotFound
		}
		return err
	}
	return nil
}
