package sapi

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func makeReq(router http.Handler, method string, url string, body io.Reader) *http.Response {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(method, url, body)
	router.ServeHTTP(w, req)
	return w.Result()
}

func TestRawStatus(t *testing.T) {
	router := gin.New()
	router.GET("/handle", Handle(func(c *gin.Context) (Res, Err) {
		return RawStatus(http.StatusOK), nil
	}))

	r := makeReq(router, "GET", "/handle", nil)
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	assert.Equal(t, http.StatusOK, r.StatusCode)
	assert.Equal(t, int64(0), r.ContentLength)
	assert.Len(t, body, 0)
}

func TestStatus(t *testing.T) {
	router := gin.New()
	router.GET("/handle", Handle(func(c *gin.Context) (Res, Err) {
		return Status(http.StatusOK), nil
	}))

	r := makeReq(router, "GET", "/handle", nil)
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	assert.Equal(t, http.StatusOK, r.StatusCode)
	assert.JSONEq(t, `{"status": 200}`, string(body))
}

func TestData_Object(t *testing.T) {
	router := gin.New()
	router.GET("/handle", Handle(func(c *gin.Context) (Res, Err) {
		return Data(http.StatusOK, gin.H{
			"field1": 123,
			"field2": true,
		}), nil
	}))

	r := makeReq(router, "GET", "/handle", nil)
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	assert.Equal(t, http.StatusOK, r.StatusCode)
	assert.JSONEq(t, `{"status": 200, "data": {"field1": 123, "field2": true}}`, string(body))
}

func TestData_Array(t *testing.T) {
	router := gin.New()
	router.GET("/handle", Handle(func(c *gin.Context) (Res, Err) {
		return Data(http.StatusOK, []gin.H{{
			"field1": 123,
			"field2": true,
		}, {
			"field1": 456,
			"field2": false,
		}}), nil
	}))

	r := makeReq(router, "GET", "/handle", nil)
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	assert.Equal(t, http.StatusOK, r.StatusCode)
	assert.JSONEq(t, `{"status": 200, "data": [{"field1": 123, "field2": true}, {"field1": 456, "field2": false}]}`, string(body))
}

func TestRedirect(t *testing.T) {
	router := gin.New()
	router.GET("/handle", Handle(func(c *gin.Context) (Res, Err) {
		return Redirect(http.StatusFound, "/location"), nil
	}))

	r := makeReq(router, "GET", "/handle", nil)
	r.Body.Close()

	assert.Equal(t, http.StatusFound, r.StatusCode)
	loc, err := r.Location()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "/location", loc.String())
}

func TestRawErrorStatus(t *testing.T) {
	router := gin.New()
	router.GET("/handle", Handle(func(c *gin.Context) (Res, Err) {
		return nil, RawErrorStatus(http.StatusBadRequest)
	}))

	r := makeReq(router, "GET", "/handle", nil)
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	assert.Equal(t, http.StatusBadRequest, r.StatusCode)
	assert.Equal(t, int64(0), r.ContentLength)
	assert.Len(t, body, 0)
}

func TestErrorStatus(t *testing.T) {
	router := gin.New()
	router.GET("/handle", Handle(func(c *gin.Context) (Res, Err) {
		return nil, ErrorStatus(http.StatusBadRequest)
	}))

	r := makeReq(router, "GET", "/handle", nil)
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	assert.Equal(t, http.StatusBadRequest, r.StatusCode)
	assert.JSONEq(t, `{"status": 400, "error": "Bad Request"}`, string(body))
}

func TestError(t *testing.T) {
	router := gin.New()
	router.GET("/handle", Handle(func(c *gin.Context) (Res, Err) {
		return nil, Error(http.StatusBadRequest, "You can't do that")
	}))

	r := makeReq(router, "GET", "/handle", nil)
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	assert.Equal(t, http.StatusBadRequest, r.StatusCode)
	assert.JSONEq(t, `{"status": 400, "error": "You can't do that"}`, string(body))
}

func TestErrorMessage(t *testing.T) {
	router := gin.New()
	router.GET("/handle", Handle(func(c *gin.Context) (Res, Err) {
		return nil, ErrorMessage(http.StatusBadRequest, "You can't do that", "For some reason")
	}))

	r := makeReq(router, "GET", "/handle", nil)
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	assert.Equal(t, http.StatusBadRequest, r.StatusCode)
	assert.JSONEq(t, `{"status": 400, "error": "You can't do that", "message": "For some reason"}`, string(body))
}

func TestAlreadyHandled(t *testing.T) {
	router := gin.New()
	router.GET("/handle", Handle(func(c *gin.Context) (Res, Err) {
		c.Status(http.StatusOK)
		return AlreadyHandled, nil
	}))

	r := makeReq(router, "GET", "/handle", nil)
	body, _ := io.ReadAll(r.Body)
	r.Body.Close()

	assert.Equal(t, http.StatusOK, r.StatusCode)
	assert.Len(t, body, 0)
}
