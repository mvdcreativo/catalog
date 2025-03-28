package file_validator

import (
	"fmt"
	"mime/multipart"
)

func ValidateFile(header *multipart.FileHeader, maxSizeMB int64, allowedTypes []string) error {

	maxSize := maxSizeMB * 1024 * 1024
	headerSizeMB := header.Size / 1024 / 1024
	if header.Size > maxSize {
		return fmt.Errorf("file too large (%d MB), max allowed is %d MB", headerSizeMB, maxSizeMB)
	}

	contentType := header.Header.Get("Content-Type")
	valid := false
	for _, t := range allowedTypes {
		if t == contentType {
			valid = true
			break
		}
	}
	if !valid {
		return fmt.Errorf("invalid content type: %s", contentType)
	}

	return nil
}
