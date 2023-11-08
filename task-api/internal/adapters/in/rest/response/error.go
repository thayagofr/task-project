package response

import "time"

type Error struct {
	Path      string    `json:"path"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
}

func NewError(path string, message string) *Error {
	return &Error{
		Timestamp: time.Now(),
		Message:   message,
		Path:      path,
	}
}
