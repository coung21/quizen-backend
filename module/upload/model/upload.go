package model

import (
	"errors"
	"quizen/common"
)

type Upload struct {
	common.SQLModel
	common.Image
}

var (
	ErrFileTooLarge   = errors.New("file too large")
	ErrFileIsNotImage = errors.New("file is not image")
	ErrCannotSaveFile = errors.New("cannot save uploaded file")
)
