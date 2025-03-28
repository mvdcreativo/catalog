package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/interfaces/storage"
)

type MinioLister struct {
	*MinioClient
}

func NewLister(client *MinioClient) storage.ListService {
	return &MinioLister{client}
}

func (m *MinioLister) ListObjects(prefix string) ([]string, error) {
	ctx := context.Background()
	var results []string

	for obj := range m.Client.ListObjects(ctx, m.BucketName, minio.ListObjectsOptions{Prefix: prefix, Recursive: true}) {
		if obj.Err != nil {
			return nil, obj.Err
		}
		results = append(results, obj.Key)
	}
	return results, nil
}
