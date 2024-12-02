package types

import "errors"

var (
	ErrInvalidStorageType = errors.New("invalid storage type")
	ErrObjectNotFound     = errors.New("object not found")
	ErrUploadFailed       = errors.New("upload failed")
)
