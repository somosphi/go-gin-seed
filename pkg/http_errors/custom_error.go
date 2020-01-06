package http_error

import (
	"fmt"
)

type (
	// CustomError struct
	//
	// * Status: HTTP Status of the error
	//
	// * Code of the error. Upper case format: EXAMPLE_ERROR
	//
	// * Title for the error. Main massage.
	//
	// * Detail is the list of details for the error
	CustomError struct {
		Status int      `json:"status"`
		Code   string   `json:"code"`
		Title  string   `json:"title"`
		Detail []string `json:"detail,omitempty"`
	}
)

func customError(status int, code, title string, detail []string) *CustomError {
	return &CustomError{status, code, title, detail}
}

// NewCustomError creates a new custom error from the givem variables
//
// * status: int -> HTTP Status code for the error
//
// * code: string -> String that represents the error
//
// * title: string -> The title of the error
//
// * details: []string -> List of details to increment the error
func NewCustomError(status int, code, title string, detail []string) error {
	return &CustomError{status, code, title, detail}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Title)
}
