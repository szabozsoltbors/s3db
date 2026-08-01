package s3

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Repository struct {
	client *s3.Client
	bucket string
}

func NewRepository(client *s3.Client, bucket string) *Repository {
	return &Repository{
		client: client,
		bucket: bucket,
	}
}

func (r *Repository) Upload(ctx context.Context, key string, data []byte) error {
	_, err := r.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &r.bucket,
		Key:    &key,
		Body:   bytes.NewReader(data),
	})

	return err
}

func (r *Repository) Delete(ctx context.Context, key string) error {
	_, err := r.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: &r.bucket,
		Key:    &key,
	})

	return err
}

func (r *Repository) List(ctx context.Context, prefix string) ([]string, error) {
	output, err := r.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: &r.bucket,
		Prefix: &prefix,
	})
	if err != nil {
		return nil, err
	}

	keys := make([]string, len(output.Contents))
	for i, object := range output.Contents {
		keys[i] = *object.Key
	}

	return keys, nil
}
