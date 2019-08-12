package roc

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	ErrMiss        = Error("cache miss")
	ErrorBucketNum = Error("bucket num must be the power of two")
)
