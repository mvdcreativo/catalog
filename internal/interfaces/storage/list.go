package storage

type ListService interface {
	ListObjects(prefix string) ([]string, error)
}
