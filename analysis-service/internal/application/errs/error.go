package errs

import "errors"

var (
	InvalidRequest = errors.New("invalid request")
	// ErrFileNotFound   = errors.New("file not found")
	// FailedToReadFile  = errors.New("failed to read file")
	// InvalidFileFormat = errors.New("only .txt files are allowed")
	// InvalidID         = errors.New("invalid id")
	// FileNotFound      = errors.New("file not found")
)
