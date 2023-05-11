package sapi

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"sync/atomic"
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

func TestHandleWithInspection(t *testing.T) {
	var n atomic.Int32

	router := gin.New()

	// Inspecting good responses
	router.GET("/handle1", HandleWithInspection(func(c *gin.Context) (Res, Err) {
		return Status(http.StatusOK), nil
	}, func(c *gin.Context, res Res) {
		n.Add(1)
		assert.Equal(t, "/handle1", c.FullPath())
		assert.Equal(t, http.StatusOK, (res.(*dataResponse)).Status)
	}, func(c *gin.Context, err Err) {
		assert.FailNow(t, "Expected an OK response")
	}))

	// Inspecting error responses
	router.GET("/handle2", HandleWithInspection(func(c *gin.Context) (Res, Err) {
		return nil, ErrorStatus(http.StatusBadRequest)
	}, func(c *gin.Context, res Res) {
		assert.FailNow(t, "Expected a Bad Request response")
	}, func(c *gin.Context, err Err) {
		n.Add(1)
		assert.Equal(t, "/handle2", c.FullPath())
		assert.Equal(t, http.StatusBadRequest, (err.(*errorResponse)).Status)
	}))

	// Invoke good response
	r := makeReq(router, "GET", "/handle1", nil)
	r.Body.Close()
	assert.Equal(t, http.StatusOK, r.StatusCode)

	// Invoke error response
	r = makeReq(router, "GET", "/handle2", nil)
	r.Body.Close()
	assert.Equal(t, http.StatusBadRequest, r.StatusCode)

	// Confirm that both calls resulted in the callbacks (and thus their asserts) being called
	assert.EqualValues(t, 2, n.Load(), "Expected both callbacks to have been invoked")
}
