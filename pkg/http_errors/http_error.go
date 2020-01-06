package http_error

import (
	"errors"
)

func New(errorType ErrorType, title string, detail []string) error {
	var creationFunc func(string, []string) error
	switch errorType {
	case BadRequestError:
		creationFunc = NewBadRequest
	case UnauthorizedError:
		creationFunc = NewUnauthorized
	case ForbiddenError:
		creationFunc = NewForbidden
	case NotFoundError:
		creationFunc = NewNotFound
	case MethodNotAllowedError:
		creationFunc = NewMethodNotAllowed
	case ConflictError:
		creationFunc = NewConflict
	case NotImplementedError:
		creationFunc = NewNotImplemented
	case InternalServerErrorError:
		creationFunc = NewInternalServerError
	default:
		return errors.New("invalid errorType")
	}

	return creationFunc(title, detail)
}
