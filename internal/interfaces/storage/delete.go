package storage

import "context"

type DeleteService interface {
	DeleteObject(crtx context.Context, objectName string) error
}
