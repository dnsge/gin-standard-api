package sapitest

import (
	sapi "github.com/dnsge/gin-standard-api"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetResStatus(t *testing.T) {
	res := sapi.RawStatus(http.StatusCreated)
	assert.Equal(t, http.StatusCreated, GetResStatus(res))
	res = sapi.Status(http.StatusCreated)
	assert.Equal(t, http.StatusCreated, GetResStatus(res))
}

func TestGetResBody(t *testing.T) {
	res := sapi.Status(http.StatusCreated)
	assert.JSONEq(t, `{"status": 201}`, GetResBody(res))
}

func TestGetResData(t *testing.T) {
	// Valid responses to get data from
	res := sapi.Data(http.StatusOK, "Hello, world")
	assert.Equal(t, "Hello, world", GetResData(res))
	res = sapi.Data(http.StatusOK, 123)
	assert.Equal(t, 123, GetResData(res))
	res = sapi.Data(http.StatusOK, []int{1, 2, 3})
	assert.Equal(t, []int{1, 2, 3}, GetResData(res))

	// Responses not created by sapi.Data return nil
	res = sapi.Status(http.StatusOK)
	assert.Equal(t, nil, GetResData(res))
}
