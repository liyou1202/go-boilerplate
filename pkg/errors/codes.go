package errors

// Success
const (
	CodeSuccess = 0
)

// Application errors (1-999)
const (
	ErrInternalError = 1
	ErrUnauthorized  = 2
)

// Database errors (1000-1099)
const (
	ErrDatabaseConnectionFailed = 1001
	ErrDatabaseQueryFailed      = 1002
	ErrDatabaseWriteFailed      = 1003
	ErrDatabaseReadFailed       = 1004
)

// AWS services errors (2000-2099)
const (
	ErrS3UploadFailed   = 2001
	ErrS3DownloadFailed = 2002
)

// External API errors (2100-2199)
const (
	ErrExternalAPI        = 2100
	ErrExternalAPITimeout = 2101
)
