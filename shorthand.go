package sapi

import "net/http"

// BadRequest is shorthand for ErrorStatus(http.StatusBadRequest)
func BadRequest() Err {
	return ErrorStatus(http.StatusBadRequest)
}

// Unauthorized is shorthand for ErrorStatus(http.StatusUnauthorized)
func Unauthorized() Err {
	return ErrorStatus(http.StatusUnauthorized)
}

// Forbidden is shorthand for ErrorStatus(http.StatusForbidden)
func Forbidden() Err {
	return ErrorStatus(http.StatusForbidden)
}

// NotFound is shorthand for ErrorStatus(http.StatusNotFound)
func NotFound() Err {
	return ErrorStatus(http.StatusNotFound)
}

// Conflict is shorthand for ErrorStatus(http.StatusConflict)
func Conflict() Err {
	return ErrorStatus(http.StatusConflict)
}

// Gone is shorthand for ErrorStatus(http.StatusGone)
func Gone() Err {
	return ErrorStatus(http.StatusGone)
}

// InternalServerError is shorthand for ErrorStatus(http.StatusInternalServerError)
func InternalServerError() Err {
	return ErrorStatus(http.StatusInternalServerError)
}
