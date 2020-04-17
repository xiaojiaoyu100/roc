package roc

// Error implements an error interface.
type Error string

// Error returns a error string.
func (e Error) Error() string {
	return string(e)
}

const (
	// ErrMiss represents a cache miss.
	ErrMiss = Error("cache miss")
)
