package sapi

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Res represents a response to a successful request.
type Res interface {
	WriteResponse(c *gin.Context)
}

// Err represents a response to an unsuccessful request.
type Err interface {
	WriteError(c *gin.Context)
}

// RawStatus writes the given status with no JSON body.
func RawStatus(status int) Res {
	return rawStatusResponse(status)
}

// Status writes the given status along with the standard API response body.
//
// Example Response:
//
//	200 OK
//
//	{
//	  "status": 200
//	}
func Status(status int) Res {
	return &dataResponse{
		Status: status,
		Data:   nil,
	}
}

// Data writes the given status along with the standard API response body
// containing the status and the data.
//
// Example Response:
//
//	200 OK
//
//	{
//	  "status": 200,
//	  "data": {"hello": "world"}
//	}
func Data(status int, data any) Res {
	return &dataResponse{
		Status: status,
		Data:   data,
	}
}

// String writes the given status and a *plain* string response. Equivalent to
// gin.String.
func String(status int, text string) Res {
	return &stringResponse{
		Status: status,
		String: text,
	}
}

// Redirect writes the given status code and sets the Location header. Equivalent
// to gin.Redirect.
func Redirect(status int, location string) Res {
	return &redirectResponse{
		Status:   status,
		Location: location,
	}
}

// RawErrorStatus writes the given status with no JSON body.
func RawErrorStatus(status int) Err {
	return rawErrorStatusResponse(status)
}

// ErrorStatus writes the given status along with the standard API response body,
// including the status text for the given status code (as defined by http.StatusText).
//
// Example Response:
//
//	400 Bad Request
//
//	{
//	  "status": 400,
//	  "error": "Bad Request"
//	}
func ErrorStatus(status int) Err {
	return Error(status, http.StatusText(status))
}

// Error writes the given status along with the standard API response body containing
// the given error text.
//
// Example Response:
//
//	403 Forbidden
//
//	{
//	  "status": 403,
//	  "error": "Sorry, you can't do that"
//	}
func Error(status int, error string) Err {
	return ErrorMessage(status, error, "")
}

// ErrorMessage writes the given status along with the standard API response body containing
// the given error text and message.
//
// Example Response:
//
//	404 Not Found
//
//	{
//	  "status": 404,
//	  "error": "User Not Found",
//	  "message": "The user you requested does not exist or is permanently banned."
//	}
func ErrorMessage(status int, error string, message string) Err {
	return &errorResponse{
		Status:  status,
		Error:   error,
		Message: message,
	}
}

// StringError writes the given status and a *plain* string response. Equivalent to
// gin.String.
func StringError(status int, text string) Err {
	return &stringResponse{
		Status: status,
		String: text,
	}
}
