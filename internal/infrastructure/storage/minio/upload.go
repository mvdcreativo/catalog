package minio

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/interfaces/storage"
)

type MinioUploader struct {
	*MinioClient
}

func NewUploader(client *MinioClient) storage.UploadService {
	return &MinioUploader{client}
}

func (m *MinioUploader) Upload(ctx context.Context, file multipart.File, header *multipart.FileHeader, refId string) (storage.FileObject, error) {
	defer file.Close()
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, file)
	if err != nil {
		return storage.FileObject{}, err
	}
	// 👇 Generamos un ID único
	imageID := uuid.New().String()
	objectName := imageID
	contentType := header.Header.Get("Content-Type")

	uploadInfo, err := m.Client.PutObject(
		ctx,
		m.BucketName,
		objectName,
		bytes.NewReader(buf.Bytes()),
		int64(buf.Len()),
		minio.PutObjectOptions{
			ContentType: contentType,
			UserMetadata: map[string]string{
				"name":  header.Filename,
				"refId": refId,
			},
			ContentEncoding: header.Header.Get("Content-Encoding"),
			ContentLanguage: header.Header.Get("Content-Language"),
		},
	)
	if err != nil {
		log.Panic(err)
		return storage.FileObject{}, err
	}

	return storage.FileObject{
		ID:          imageID,
		URL:         fmt.Sprintf("%s/%s/%s", m.BaseURL, m.BucketName, objectName),
		FileName:    header.Filename,
		Size:        uploadInfo.Size,
		ETag:        uploadInfo.ETag,
		ContentType: contentType,
		RefId:       refId,
	}, nil
}
