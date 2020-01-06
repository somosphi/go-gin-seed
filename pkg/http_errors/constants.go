package http_error

const (
	BAD_REQUEST_CODE        = "BAD_REQUEST"        // 400
	UNAUTHORIZED_CODE       = "UNAUTHORIZED"       // 401
	FORBIDDEN_CODE          = "FORBIDDEN"          // 403
	NOT_FOUND_CODE          = "NOT_FOUND"          // 404
	METHOD_NOT_ALLOWED_CODE = "METHOD_NOT_ALLOWED" // 405
	CONFLICT_CODE           = "CONFLICT"           // 409

	INTERNAL_SERVER_ERROR_CODE = "INTERNAL_SERVER_ERROR" // 500
	NOT_IMPLEMENTED_CODE       = "NOT_IMPLEMENTED"       // 501
)

type ErrorType int

const (
	BadRequestError ErrorType = iota + 1
	UnauthorizedError
	ForbiddenError
	NotFoundError
	MethodNotAllowedError
	ConflictError
	InternalServerErrorError
	NotImplementedError
)
