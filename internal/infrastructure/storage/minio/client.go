package minio

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient struct {
	Client     *minio.Client
	BucketName string
	BaseURL    string
}

func NewMinioClient(endpoint, accessKey, secretKey, bucket, baseURL string, useSSL bool) *MinioClient {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("❌ MinIO client init failed: %v", err)
	}

	// Crear bucket si no existe
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucket)
	if err != nil {
		log.Fatalf("❌ Bucket check failed: %v", err)
	}
	if !exists {
		if err := client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{}); err != nil {
			log.Fatalf("❌ Bucket creation failed: %v", err)
		}
	}

	log.Println("✅ Conectado a Minio")
	return &MinioClient{
		Client:     client,
		BucketName: bucket,
		BaseURL:    baseURL,
	}
}

// // Upload sube un archivo representado en []byte a MinIO y devuelve la URL pública.
// // Se asume que el bucket es público o que se genera la URL de acceso.
// func (c *Client) Upload(ctx context.Context, objectName string, data []byte, contentType string) (string, error) {
// 	reader := bytes.NewReader(data)
// 	_, err := c.Minio.PutObject(ctx, c.BucketName, objectName, reader, int64(len(data)), minio.PutObjectOptions{
// 		ContentType: contentType,
// 	})
// 	if err != nil {
// 		return "", fmt.Errorf("error al subir objeto: %v", err)
// 	}

// 	// Construir la URL pública (ajusta según tu configuración de red o CDN).
// 	publicURL := fmt.Sprintf("http://%s/%s/%s", c.Endpoint, c.BucketName, objectName)
// 	return publicURL, nil
// }
