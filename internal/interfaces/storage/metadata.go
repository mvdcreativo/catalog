package storage

type MetadataService interface {
	GetMetadata(objectName string) (map[string]string, error)
}
