package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/interfaces/storage"
)

type MinioMetaFetcher struct {
	*MinioClient
}

func NewMetaFetcher(client *MinioClient) storage.MetadataService {
	return &MinioMetaFetcher{client}
}

func (m *MinioMetaFetcher) GetMetadata(objectName string) (map[string]string, error) {
	info, err := m.Client.StatObject(context.Background(), m.BucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		return nil, err
	}
	return info.UserMetadata, nil
}
