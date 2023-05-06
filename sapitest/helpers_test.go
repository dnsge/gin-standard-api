package sapitest

import (
	"github.com/dnsge/gin-standard-api"
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

func TestGetErrStatus(t *testing.T) {
	err := sapi.RawErrorStatus(http.StatusInternalServerError)
	assert.Equal(t, http.StatusInternalServerError, GetErrStatus(err))
	err = sapi.ErrorStatus(http.StatusInternalServerError)
	assert.Equal(t, http.StatusInternalServerError, GetErrStatus(err))
}

func TestGetErrBody(t *testing.T) {
	err := sapi.Error(http.StatusInternalServerError, "Oh no")
	assert.JSONEq(t, `{"status": 500, "error": "Oh no"}`, GetErrBody(err))
}
