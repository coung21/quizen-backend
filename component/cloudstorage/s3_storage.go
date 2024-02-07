package cloudstorage

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"quizen/common"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Storage struct {
	bucketName string
	region     string
	accessKey  string
	secretKey  string
	domain     string
	session    *session.Session
}

func NewS3Storage(bucketName, region, accessKey, secretKey, domain string) *s3Storage {
	provider := &s3Storage{
		bucketName: bucketName,
		region:     region,
		accessKey:  accessKey,
		secretKey:  secretKey,
		domain:     domain,
	}

	s3Session, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})

	if err != nil {
		panic(err)
	}

	provider.session = s3Session

	return provider
}

// SaveUploadedFile saves the uploaded file to S3
func (s *s3Storage) SaveUploadedFile(ctx context.Context, data []byte, dst string) (*common.Image, error) {
	fileBytes := bytes.NewReader(data)
	fileTypes := http.DetectContentType(data)

	_, err := s3.New(s.session).PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(dst),
		Body:        fileBytes,
		ContentType: aws.String(fileTypes),
		ACL:         aws.String("private"),
	})

	if err != nil {
		return nil, err
	}

	img := &common.Image{
		Url:       fmt.Sprintf("%s%s", s.domain, dst),
		CloudName: "s3",
	}

	return img, nil
}
