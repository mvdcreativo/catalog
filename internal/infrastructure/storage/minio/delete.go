package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/interfaces/storage"
)

type MinioDeleter struct {
	*MinioClient
}

func NewDeleter(client *MinioClient) storage.DeleteService {
	return &MinioDeleter{client}
}

func (m *MinioDeleter) DeleteObject(ctx context.Context, objectName string) error {

	return m.Client.RemoveObject(ctx, m.BucketName, objectName, minio.RemoveObjectOptions{})
}
