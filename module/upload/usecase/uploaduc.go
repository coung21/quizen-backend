package usecase

import (
	"bytes"
	"context"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"path/filepath"
	"quizen/common"
	"quizen/component/cloudstorage"
	"quizen/module/upload/model"
	"strings"
	"time"
)

type Usecase interface {
	Upload(ctx context.Context, file []byte, filename, folder string) (*common.Image, error)
}

type uploadUc struct {
	provider cloudstorage.UploadProvider
}

func NewUploadUc(provider cloudstorage.UploadProvider) uploadUc {
	return uploadUc{provider}
}

func (uc uploadUc) Upload(ctx context.Context, data []byte, fileName, folder string) (*common.Image, error) {
	reader := bytes.NewReader(data)
	width, height, err := getImageDimension(reader)

	if err != nil {
		return nil, model.ErrFileIsNotImage
	}

	if strings.TrimSpace(folder) == "" {
		folder = "images"
	}

	if width > 1024 || height > 1024 {
		return nil, model.ErrFileTooLarge
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := uc.provider.SaveUploadedFile(ctx, data, fileName)

	if err != nil {
		return nil, model.ErrCannotSaveFile
	}

	img.Width = width
	img.Height = height
	img.Ext = fileExt

	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)
	if err != nil {
		log.Println("err: ", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
