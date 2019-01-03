package restchain

import (
	"fmt"
)

type Error struct {
	Code    int                    `json:"code,omitempty"`
	Message string                 `json:"message,omitempty"`
	Meta    map[string]interface{} `json:"meta,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%s. Code: %d", e.Message, e.Code)
}

func NewError(code int, message string, meta map[string]interface{}) error {
	return &Error{
		Code:    code,
		Message: message,
		Meta:    meta,
	}
}
