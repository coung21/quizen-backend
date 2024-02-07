package cloudstorage

import (
	"context"
	"quizen/common"
)

type UploadProvider interface {
	SaveUploadedFile(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
