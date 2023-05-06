package sapitest

import (
	sapi "github.com/dnsge/gin-standard-api"
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

// GetResData returns the data contained in the given response. This requires
// the res to have been created via sapi.Data, otherwise nil is returned.
func GetResData(res sapi.Res) any {
	asData, ok := res.(interface {
		GetData() any
	})
	if !ok {
		return nil
	}
	return asData.GetData()
}
