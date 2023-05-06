package sapitest

import (
	"github.com/dnsge/gin-standard-api"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
)

func execRes(res sapi.Res) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	res.WriteResponse(c)
	// ref: https://github.com/gin-gonic/gin/issues/1120
	c.Writer.WriteHeaderNow()
	return recorder
}

func execErr(err sapi.Err) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	err.WriteError(c)
	// ref: https://github.com/gin-gonic/gin/issues/1120
	c.Writer.WriteHeaderNow()
	return recorder
}

// GetResStatus returns the status code contained within the given response.
// This method should be exclusively used for testing.
func GetResStatus(res sapi.Res) int {
	recorder := execRes(res)
	return recorder.Code
}

// GetResHeaders returns the headers written for the given response.
// This method should be exclusively used for testing.
func GetResHeaders(res sapi.Res) http.Header {
	recorder := execRes(res)
	return recorder.Header()
}

// GetResBody returns the body written for the given response.
// This method should be exclusively used for testing.
func GetResBody(res sapi.Res) string {
	recorder := execRes(res)
	return recorder.Body.String()
}

// GetErrStatus returns the status code contained within the given error response.
// This method should be exclusively used for testing.
func GetErrStatus(err sapi.Err) int {
	recorder := execErr(err)
	return recorder.Code
}

// GetErrHeaders returns the headers written for the given error response.
// This method should be exclusively used for testing.
func GetErrHeaders(err sapi.Err) http.Header {
	recorder := execErr(err)
	return recorder.Header()
}

// GetErrBody returns the body written for the given error response.
// This method should be exclusively used for testing.
func GetErrBody(err sapi.Err) string {
	recorder := execErr(err)
	return recorder.Body.String()
}
