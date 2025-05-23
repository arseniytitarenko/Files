package errs

import "errors"

var (
	ErrFileNotFound  = errors.New("file not found")
	FailedToReadFile = errors.New("failed to read file")
)
