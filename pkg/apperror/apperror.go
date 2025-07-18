package apperror

import "github.com/vektah/gqlparser/v2/gqlerror"

func Internal(msg string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: msg,
		Extensions: map[string]interface{}{
			"code": "INTERNAL_SERVER_ERROR",
		},
	}
}

func Unauthorized(msg string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: msg,
		Extensions: map[string]interface{}{
			"code": "UNAUTHORIZED",
		},
	}
}

func BadRequest(msg string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: msg,
		Extensions: map[string]interface{}{
			"code": "BAD_REQUEST",
		},
	}
}

func Conflict(msg string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: msg,
		Extensions: map[string]interface{}{
			"code": "CONFLICT",
		},
	}
}
func NotFound(msg string) *gqlerror.Error {
	return &gqlerror.Error{
		Message: msg,
		Extensions: map[string]interface{}{
			"code": "NOT_FOUND",
		},
	}
}
