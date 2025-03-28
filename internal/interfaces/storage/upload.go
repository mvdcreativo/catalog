package storage

import (
	"context"
	"mime/multipart"
)

type UploadService interface {
	Upload(ctx context.Context, file multipart.File, header *multipart.FileHeader, refId string) (FileObject, error)
}
