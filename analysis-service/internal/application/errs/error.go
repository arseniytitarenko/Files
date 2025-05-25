package errs

import "errors"

var (
	InvalidRequest      = errors.New("invalid request")
	ErrLocationNotFound = errors.New("location not found")
	ExternalApiError    = errors.New("external api error")
	StorageServiceError = errors.New("storage service error")
	// FailedToReadFile  = errors.New("failed to read file")
	// InvalidFileFormat = errors.New("only .txt files are allowed")
	// InvalidID         = errors.New("invalid id")
	// FileNotFound      = errors.New("file not found")
)
