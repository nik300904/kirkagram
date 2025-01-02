package s3

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"io"
	"kirkagram/internal/storage"
)

const bucketName = "kirkagram"

type PhotoS3Storage struct {
	client *s3.Client
}

func NewUserS3Storage(client *s3.Client) *PhotoS3Storage {
	return &PhotoS3Storage{client: client}
}

func (u *PhotoS3Storage) GetPhoto(key string) ([]byte, error) {
	const op = "storage.s3.GetPhoto"

	result, err := u.client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		var nsk *types.NoSuchKey
		if errors.As(err, &nsk) {
			return nil, fmt.Errorf("%s: %w", op, storage.ErrNoSuchKey)
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer result.Body.Close()

	return io.ReadAll(result.Body)
}
