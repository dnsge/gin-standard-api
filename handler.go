package sapi

import "github.com/gin-gonic/gin"

// AlreadyHandled is a sentinel Res instance that indicates to Handle that
// the called handler has already written a response and that no interpretation
// of the Res or Err is necessary.
var AlreadyHandled = Res(nil)

// Handler is a gin.HandlerFunc which returns a Res or an Err.
type Handler func(c *gin.Context) (Res, Err)

// Handle wraps a Handler into a gin.HandlerFunc instance.
//
// Example usage:
//
//	router.GET("/api/endpoint", sapi.Handle(MyEndpoint))
func Handle(handler Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := handler(c)
		if err != nil {
			err.WriteError(c)
			return
		}

		// Assume that the handler already somehow took care of a response
		// and knows what it's doing.
		if res != nil && res != AlreadyHandled {
			res.WriteResponse(c)
		}
	}
}

type BeforeResFunc = func(c *gin.Context, res Res)
type BeforeErrorFunc = func(c *gin.Context, err Err)

// HandleWithInspection wraps a Handler into a gin.HandlerFunc instance.
// The inspection callbacks are invoked before writing the response and writing
// the error, respectively.
func HandleWithInspection(handler Handler, beforeRes BeforeResFunc, beforeError BeforeErrorFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := handler(c)
		if err != nil {
			if beforeError != nil {
				beforeError(c, err)
			}
			err.WriteError(c)
			return
		}

		// Assume that the handler already somehow took care of a response
		// and knows what it's doing.
		if res != nil && res != AlreadyHandled {
			if beforeRes != nil {
				beforeRes(c, res)
			}
			res.WriteResponse(c)
		}
	}
}
