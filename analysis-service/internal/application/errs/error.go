package errs

import "errors"

var (
	InvalidRequest      = errors.New("invalid request")
	ErrLocationNotFound = errors.New("location not found")
	ExternalApiError    = errors.New("external api error")
	StorageServiceError = errors.New("storage service error")
	InvalidID           = errors.New("invalid id")
)
