package http_error

import (
	"fmt"
	"net/http"
)

type (
	// BadRequest error struct
	BadRequest struct{ *CustomError }

	// Unauthorized error struct
	Unauthorized struct{ *CustomError }

	// Forbidden error struct
	Forbidden struct{ *CustomError }

	// NotFound error struct
	NotFound struct{ *CustomError }

	// MethodNotAllowed error struct
	MethodNotAllowed struct{ *CustomError }

	// Conflict error struct
	Conflict struct{ *CustomError }

	// InternalServerError error struct
	InternalServerError struct{ *CustomError }

	// NotImplemented error struct
	NotImplemented struct{ *CustomError }
)

// ------------------ Bad Request ---------------------- //

// Create a new Bad request error
//
// * Status Code: 400
//
// * Error Code: BAD_REQUEST
//
// Example: `BAD_REQUEST: you executed some bad request ;)`
func NewBadRequest(title string, detail []string) error {
	cer := customError(http.StatusBadRequest, BAD_REQUEST_CODE, title, detail)
	return &BadRequest{cer}
}

// String representation of a Bad request error
func (e *BadRequest) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Title)
}

// ------------------ Unauthorized ---------------------- //

// Create a new Unauthorized error
//
// * Status Code: 401
//
// * Error Code: UNAUTHORIZED
//
// Example: `UNAUTHORIZED: you doesn't have authorization for some operation`
func NewUnauthorized(title string, detail []string) error {
	cer := customError(http.StatusUnauthorized, UNAUTHORIZED_CODE, title, detail)
	return &Unauthorized{cer}
}

// String representation of a Unauthorized error
func (e *Unauthorized) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Title)
}

// ------------------ Forbidden ---------------------- //

// Create a new Forbidden error
//
// * Status Code: 403
//
// * Error Code: FORBIDDEN
//
// Example: `FORBIDDEN: some resource is forbidden`
func NewForbidden(title string, detail []string) error {
	cer := customError(http.StatusForbidden, FORBIDDEN_CODE, title, detail)
	return &Forbidden{cer}
}

// String representation of a Forbidden error
func (e *Forbidden) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Title)
}

// ------------------ Not Found ---------------------- //

// Create a new NotFound error
//
// * Status Code: 404
//
// * Error Code: NOT_FOUND
//
// Example: `NOT_FOUND: some resource wasn't found`
func NewNotFound(title string, detail []string) error {
	cer := customError(http.StatusNotFound, NOT_FOUND_CODE, title, detail)
	return &NotFound{cer}
}

// String representation of a NotFound error
func (e *NotFound) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Title)
}

// -------------------  Method Not Allowed  ------------------ //

// Create a new MethodNotAllowed error
//
// * Status Code: 405
//
// * Error Code: METHOD_NOT_ALLOWED
//
// Example: `METHOD_NOT_ALLOWED: some method that isn't allowed`
func NewMethodNotAllowed(title string, detail []string) error {
	cer := customError(
		http.StatusMethodNotAllowed,
		METHOD_NOT_ALLOWED_CODE,
		title,
		detail,
	)
	return &MethodNotAllowed{cer}
}

// String representation of a MethodNotAllowed error
func (e *MethodNotAllowed) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Title)
}

// ------------------ Conflict ---------------------- //

// Create a new Conflict error
//
// * Status Code: 409
//
// * Error Code: CONFLICT
//
// Example: `CONFLICT: some conflict message`
func NewConflict(title string, detail []string) error {
	cer := customError(http.StatusConflict, CONFLICT_CODE, title, detail)
	return &Conflict{cer}
}

// String representation of a Conflict error
func (e *Conflict) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Title)
}

// ------------------ Internal Server Error ---------------------- //

// Create a new InternalServerError error
//
// * Status Code: 500
//
// * Error Code: INTERNAL_SERVER_ERROR
//
// Example: `INTERNAL_SERVER_ERROR: some unexpected error`
func NewInternalServerError(title string, detail []string) error {
	cer := customError(
		http.StatusInternalServerError,
		INTERNAL_SERVER_ERROR_CODE,
		title,
		detail,
	)
	return &InternalServerError{cer}
}

// String representation of a InternalServerError error
func (e *InternalServerError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Title)
}

// ------------------ Not Implemented ---------------------- //

// Create a new NotImplemented error
//
// * Status Code: 501
//
// * Error Code: NOT_IMPLEMENTED
//
// Example: `NOT_IMPLEMENTED: some functionality isn't implemented`
func NewNotImplemented(title string, detail []string) error {
	cer := customError(
		http.StatusNotImplemented,
		NOT_IMPLEMENTED_CODE,
		title,
		detail,
	)
	return &NotImplemented{cer}
}

// String representation of a NotImplemented error
func (e *NotImplemented) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Title)
}
